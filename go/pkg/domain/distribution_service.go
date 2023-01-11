package domain

type DistributionService string

const (
	DistributionServiceUnknown      DistributionService = "unknown"
	DistributionServiceSpotify      DistributionService = "spotify"
	DistributionServiceAppleMusic   DistributionService = "apple_music"
	DistributionServiceYouTubeMusic DistributionService = "youtube_music"
	DistributionServiceLineMusic    DistributionService = "line_music"
	DistributionServiceItunes       DistributionService = "itunes"
	DistributionServiceYouTube      DistributionService = "youtube"
	DistributionServiceNicoVideo    DistributionService = "nicovideo"
	DistributionServiceSoundCloud   DistributionService = "sound_cloud"
)

func GetDistributionService(ds string) DistributionService {
	switch ds {
	case "spotify":
		return DistributionServiceSpotify
	case "apple_music":
		return DistributionServiceAppleMusic
	case "youtube_music":
		return DistributionServiceYouTubeMusic
	case "line_music":
		return DistributionServiceLineMusic
	case "itunes":
		return DistributionServiceItunes
	case "youtube":
		return DistributionServiceYouTube
	case "nicovideo":
		return DistributionServiceNicoVideo
	case "sound_cloud":
		return DistributionServiceSoundCloud
	}
	return DistributionServiceUnknown
}
