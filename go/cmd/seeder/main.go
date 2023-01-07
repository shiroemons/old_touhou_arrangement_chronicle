package main

import (
	"context"
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/spkg/bom"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
)

type Product struct {
	bun.BaseModel `bun:"table:products,alias:p"`

	ID           string  `csv:"id" bun:",pk"`
	Name         string  `csv:"name" bun:"name,nullzero,notnull"`
	ShortName    string  `csv:"short_name" bun:"short_name,nullzero,notnull"`
	ProductType  string  `csv:"product_type" bun:"product_type,nullzero,notnull"`
	SeriesNumber float64 `csv:"series_number" bun:"series_number,nullzero,notnull"`
}

type OriginalSong struct {
	bun.BaseModel `bun:"table:original_songs,alias:os"`

	ID          string `csv:"id" bun:",pk"`
	ProductID   string `csv:"product_id" bun:"product_id,nullzero,notnull"`
	TrackNumber int    `csv:"track_number" bun:"track_number,nullzero,notnull"`
	Name        string `csv:"name" bun:"name,nullzero,notnull"`
	Composer    string `csv:"composer" bun:"composer,nullzero,notnull,default:''"`
	Arranger    string `csv:"arranger" bun:"arranger,nullzero,notnull,default:''"`
	SourceID    string `csv:"source_id" bun:"source_id,nullzero,notnull,default:''"`
	Original    bool   `csv:"is_original" bun:"is_original,notnull"`
}

func main() {
	ctx := context.Background()
	db := initDB()

	fn := func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(bom.NewReader(in)) // BOMの回避
		r.Comma = '\t'                        // 区切り文字をタブに変更
		r.Comment = '#'                       // #で始まる行はコメントと見なしスキップ
		return r
	}
	gocsv.SetCSVReader(fn)

	importProducts(ctx, db)
	importOriginalSongs(ctx, db)
}

func initDB() *bun.DB {
	config, err := pgx.ParseConfig(os.Getenv("CONNECT_URL"))
	if err != nil {
		panic(err)
	}

	sqldb := stdlib.OpenDB(*config)
	db := bun.NewDB(sqldb, pgdialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(
		// disable the hook
		bundebug.WithEnabled(false),

		// BUNDEBUG=1 logs failed queries
		// BUNDEBUG=2 logs all queries
		bundebug.FromEnv("BUNDEBUG"),
	))

	var v string
	if err = db.NewSelect().ColumnExpr("version()").Scan(context.Background(), &v); err != nil {
		log.Fatal(err)
	}
	log.Println(v)

	return db
}

func importProducts(ctx context.Context, db *bun.DB) {
	log.Println("start products import.")

	f, err := os.Open("../db/fixtures/products.tsv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var lines []Product
	if err = gocsv.UnmarshalFile(f, &lines); err != nil {
		log.Fatal(err)
	}

	_, err = db.NewInsert().Model(&lines).
		On("CONFLICT (id) DO UPDATE").
		Set("name = EXCLUDED.name").
		Set("short_name = EXCLUDED.short_name").
		Set("product_type = EXCLUDED.product_type").
		Set("series_number = EXCLUDED.series_number").
		Exec(ctx)
	if err != nil {
		panic(err)
	}

	log.Println("finish products import.")
}

func importOriginalSongs(ctx context.Context, db *bun.DB) {
	log.Println("start original_songs import.")

	f, err := os.Open("../db/fixtures/original_songs.tsv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var lines []OriginalSong
	if err = gocsv.UnmarshalFile(f, &lines); err != nil {
		log.Fatal(err)
	}

	_, err = db.NewInsert().Model(&lines).
		On("CONFLICT (id) DO UPDATE").
		Set("product_id = EXCLUDED.product_id").
		Set("name = EXCLUDED.name").
		Set("composer = EXCLUDED.composer").
		Set("arranger = EXCLUDED.arranger").
		Set("track_number = EXCLUDED.track_number").
		Set("is_original = EXCLUDED.is_original").
		Set("source_id = EXCLUDED.source_id").
		Exec(ctx)
	if err != nil {
		panic(err)
	}

	log.Println("finish original_songs import.")
}
