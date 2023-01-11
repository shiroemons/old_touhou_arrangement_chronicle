package domain

import (
	"testing"
)

func TestGetDistributionService(t *testing.T) {
	t.Parallel()
	type args struct {
		ds string
	}
	tests := []struct {
		name string
		args args
		want DistributionService
	}{
		{
			name: "Spotify",
			args: args{ds: "spotify"},
			want: DistributionServiceSpotify,
		},
		{
			name: "Apple Music",
			args: args{ds: "apple_music"},
			want: DistributionServiceAppleMusic,
		},
		{
			name: "YouTube Music",
			args: args{ds: "youtube_music"},
			want: DistributionServiceYouTubeMusic,
		},
		{
			name: "LINE MUSIC",
			args: args{ds: "line_music"},
			want: DistributionServiceLineMusic,
		},
		{
			name: "iTunes",
			args: args{ds: "itunes"},
			want: DistributionServiceItunes,
		},
		{
			name: "YouTube",
			args: args{ds: "youtube"},
			want: DistributionServiceYouTube,
		},
		{
			name: "ニコニコ動画",
			args: args{ds: "nicovideo"},
			want: DistributionServiceNicoVideo,
		},
		{
			name: "SoundCloud",
			args: args{ds: "sound_cloud"},
			want: DistributionServiceSoundCloud,
		},
		{
			name: "SoundCloud",
			args: args{ds: "sound_cloud"},
			want: DistributionServiceSoundCloud,
		},
		{
			name: "Unknown",
			args: args{ds: "unknown"},
			want: DistributionServiceUnknown,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := GetDistributionService(tt.args.ds); got != tt.want {
				t.Errorf("GetDistributionService() = %v, want %v", got, tt.want)
			}
		})
	}
}
