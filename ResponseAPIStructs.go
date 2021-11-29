package CurseAPIWrapper

import "time"

type ModsSearchResponse struct {
	Data []struct {
		ID     int    `json:"id"`
		GameID int    `json:"gameId"`
		Name   string `json:"name"`
		Slug   string `json:"slug"`
		Links  struct {
			WebsiteURL string `json:"websiteUrl"`
			WikiURL    string `json:"wikiUrl"`
			IssuesURL  string `json:"issuesUrl"`
			SourceURL  string `json:"sourceUrl"`
		} `json:"links"`
		Summary           string `json:"summary"`
		Status            int    `json:"status"`
		DownloadCount     float64    `json:"downloadCount"`
		IsFeatured        bool   `json:"isFeatured"`
		PrimaryCategoryID int    `json:"primaryCategoryId"`
		Categories        []struct {
			ID               int       `json:"id"`
			GameID           int       `json:"gameId"`
			Name             string    `json:"name"`
			Slug             string    `json:"slug"`
			URL              string    `json:"url"`
			IconURL          string    `json:"iconUrl"`
			DateModified     time.Time `json:"dateModified"`
			IsClass          bool      `json:"isClass"`
			ClassID          int       `json:"classId"`
			ParentCategoryID int       `json:"parentCategoryId"`
		} `json:"categories"`
		Authors []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"authors"`
		Logo struct {
			ID           int    `json:"id"`
			ModID        int    `json:"modId"`
			Title        string `json:"title"`
			Description  string `json:"description"`
			ThumbnailURL string `json:"thumbnailUrl"`
			URL          string `json:"url"`
		} `json:"logo"`
		Screenshots []struct {
			ID           int    `json:"id"`
			ModID        int    `json:"modId"`
			Title        string `json:"title"`
			Description  string `json:"description"`
			ThumbnailURL string `json:"thumbnailUrl"`
			URL          string `json:"url"`
		} `json:"screenshots"`
		MainFileID  int `json:"mainFileId"`
		LatestFiles []struct {
			ID          int    `json:"id"`
			GameID      int    `json:"gameId"`
			ModID       int    `json:"modId"`
			IsAvailable bool   `json:"isAvailable"`
			DisplayName string `json:"displayName"`
			FileName    string `json:"fileName"`
			ReleaseType int    `json:"releaseType"`
			FileStatus  int    `json:"fileStatus"`
			Hashes      []struct {
				Value string `json:"value"`
				Algo  int    `json:"algo"`
			} `json:"hashes"`
			FileDate             time.Time `json:"fileDate"`
			FileLength           int       `json:"fileLength"`
			DownloadCount        int       `json:"downloadCount"`
			DownloadURL          string    `json:"downloadUrl"`
			GameVersions         []string  `json:"gameVersions"`
			SortableGameVersions []struct {
				GameVersionName        string    `json:"gameVersionName"`
				GameVersionPadded      string    `json:"gameVersionPadded"`
				GameVersion            string    `json:"gameVersion"`
				GameVersionReleaseDate time.Time `json:"gameVersionReleaseDate"`
				GameVersionTypeID      int       `json:"gameVersionTypeId"`
			} `json:"sortableGameVersions"`
			Dependencies []struct {
				ModID        int `json:"modId"`
				FileID       int `json:"fileId"`
				RelationType int `json:"relationType"`
			} `json:"dependencies"`
			ExposeAsAlternative bool `json:"exposeAsAlternative"`
			ParentProjectFileID int  `json:"parentProjectFileId"`
			AlternateFileID     int  `json:"alternateFileId"`
			IsServerPack        bool `json:"isServerPack"`
			ServerPackFileID    int  `json:"serverPackFileId"`
			FileFingerprint     int  `json:"fileFingerprint"`
			Modules             []struct {
				Name        string `json:"name"`
				Fingerprint int    `json:"fingerprint"`
			} `json:"modules"`
		} `json:"latestFiles"`
		LatestFilesIndexes []struct {
			GameVersion       string `json:"gameVersion"`
			FileID            int    `json:"fileId"`
			Filename          string `json:"filename"`
			ReleaseType       int    `json:"releaseType"`
			GameVersionTypeID int    `json:"gameVersionTypeId"`
			ModLoader         int    `json:"modLoader"`
		} `json:"latestFilesIndexes"`
		DateCreated  time.Time `json:"dateCreated"`
		DateModified time.Time `json:"dateModified"`
		DateReleased time.Time `json:"dateReleased"`
	} `json:"data"`
	Pagination struct {
		Index       int `json:"index"`
		PageSize    int `json:"pageSize"`
		ResultCount int `json:"resultCount"`
		TotalCount  int `json:"totalCount"`
	} `json:"pagination"`
}

type GetModResults struct {
	Data struct {
		ID     int    `json:"id"`
		GameID int    `json:"gameId"`
		Name   string `json:"name"`
		Slug   string `json:"slug"`
		Links  struct {
			WebsiteURL string `json:"websiteUrl"`
			WikiURL    string `json:"wikiUrl"`
			IssuesURL  string `json:"issuesUrl"`
			SourceURL  string `json:"sourceUrl"`
		} `json:"links"`
		Summary           string `json:"summary"`
		Status            int    `json:"status"`
		DownloadCount     int    `json:"downloadCount"`
		IsFeatured        bool   `json:"isFeatured"`
		PrimaryCategoryID int    `json:"primaryCategoryId"`
		Categories        []struct {
			ID               int       `json:"id"`
			GameID           int       `json:"gameId"`
			Name             string    `json:"name"`
			Slug             string    `json:"slug"`
			URL              string    `json:"url"`
			IconURL          string    `json:"iconUrl"`
			DateModified     time.Time `json:"dateModified"`
			IsClass          bool      `json:"isClass"`
			ClassID          int       `json:"classId"`
			ParentCategoryID int       `json:"parentCategoryId"`
		} `json:"categories"`
		Authors []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"authors"`
		Logo struct {
			ID           int    `json:"id"`
			ModID        int    `json:"modId"`
			Title        string `json:"title"`
			Description  string `json:"description"`
			ThumbnailURL string `json:"thumbnailUrl"`
			URL          string `json:"url"`
		} `json:"logo"`
		Screenshots []struct {
			ID           int    `json:"id"`
			ModID        int    `json:"modId"`
			Title        string `json:"title"`
			Description  string `json:"description"`
			ThumbnailURL string `json:"thumbnailUrl"`
			URL          string `json:"url"`
		} `json:"screenshots"`
		MainFileID  int `json:"mainFileId"`
		LatestFiles []struct {
			ID          int    `json:"id"`
			GameID      int    `json:"gameId"`
			ModID       int    `json:"modId"`
			IsAvailable bool   `json:"isAvailable"`
			DisplayName string `json:"displayName"`
			FileName    string `json:"fileName"`
			ReleaseType int    `json:"releaseType"`
			FileStatus  int    `json:"fileStatus"`
			Hashes      []struct {
				Value string `json:"value"`
				Algo  int    `json:"algo"`
			} `json:"hashes"`
			FileDate             time.Time `json:"fileDate"`
			FileLength           int       `json:"fileLength"`
			DownloadCount        int       `json:"downloadCount"`
			DownloadURL          string    `json:"downloadUrl"`
			GameVersions         []string  `json:"gameVersions"`
			SortableGameVersions []struct {
				GameVersionName        string    `json:"gameVersionName"`
				GameVersionPadded      string    `json:"gameVersionPadded"`
				GameVersion            string    `json:"gameVersion"`
				GameVersionReleaseDate time.Time `json:"gameVersionReleaseDate"`
				GameVersionTypeID      int       `json:"gameVersionTypeId"`
			} `json:"sortableGameVersions"`
			Dependencies []struct {
				ModID        int `json:"modId"`
				FileID       int `json:"fileId"`
				RelationType int `json:"relationType"`
			} `json:"dependencies"`
			ExposeAsAlternative bool `json:"exposeAsAlternative"`
			ParentProjectFileID int  `json:"parentProjectFileId"`
			AlternateFileID     int  `json:"alternateFileId"`
			IsServerPack        bool `json:"isServerPack"`
			ServerPackFileID    int  `json:"serverPackFileId"`
			FileFingerprint     int  `json:"fileFingerprint"`
			Modules             []struct {
				Name        string `json:"name"`
				Fingerprint int    `json:"fingerprint"`
			} `json:"modules"`
		} `json:"latestFiles"`
		LatestFilesIndexes []struct {
			GameVersion       string `json:"gameVersion"`
			FileID            int    `json:"fileId"`
			Filename          string `json:"filename"`
			ReleaseType       int    `json:"releaseType"`
			GameVersionTypeID int    `json:"gameVersionTypeId"`
			ModLoader         int    `json:"modLoader"`
		} `json:"latestFilesIndexes"`
		DateCreated  time.Time `json:"dateCreated"`
		DateModified time.Time `json:"dateModified"`
		DateReleased time.Time `json:"dateReleased"`
	} `json:"data"`
}

type GetModsResults struct {
	Data []struct {
		ID     int    `json:"id"`
		GameID int    `json:"gameId"`
		Name   string `json:"name"`
		Slug   string `json:"slug"`
		Links  struct {
			WebsiteURL string `json:"websiteUrl"`
			WikiURL    string `json:"wikiUrl"`
			IssuesURL  string `json:"issuesUrl"`
			SourceURL  string `json:"sourceUrl"`
		} `json:"links"`
		Summary           string `json:"summary"`
		Status            int    `json:"status"`
		DownloadCount     int    `json:"downloadCount"`
		IsFeatured        bool   `json:"isFeatured"`
		PrimaryCategoryID int    `json:"primaryCategoryId"`
		Categories        []struct {
			ID               int       `json:"id"`
			GameID           int       `json:"gameId"`
			Name             string    `json:"name"`
			Slug             string    `json:"slug"`
			URL              string    `json:"url"`
			IconURL          string    `json:"iconUrl"`
			DateModified     time.Time `json:"dateModified"`
			IsClass          bool      `json:"isClass"`
			ClassID          int       `json:"classId"`
			ParentCategoryID int       `json:"parentCategoryId"`
		} `json:"categories"`
		Authors []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"authors"`
		Logo struct {
			ID           int    `json:"id"`
			ModID        int    `json:"modId"`
			Title        string `json:"title"`
			Description  string `json:"description"`
			ThumbnailURL string `json:"thumbnailUrl"`
			URL          string `json:"url"`
		} `json:"logo"`
		Screenshots []struct {
			ID           int    `json:"id"`
			ModID        int    `json:"modId"`
			Title        string `json:"title"`
			Description  string `json:"description"`
			ThumbnailURL string `json:"thumbnailUrl"`
			URL          string `json:"url"`
		} `json:"screenshots"`
		MainFileID  int `json:"mainFileId"`
		LatestFiles []struct {
			ID          int    `json:"id"`
			GameID      int    `json:"gameId"`
			ModID       int    `json:"modId"`
			IsAvailable bool   `json:"isAvailable"`
			DisplayName string `json:"displayName"`
			FileName    string `json:"fileName"`
			ReleaseType int    `json:"releaseType"`
			FileStatus  int    `json:"fileStatus"`
			Hashes      []struct {
				Value string `json:"value"`
				Algo  int    `json:"algo"`
			} `json:"hashes"`
			FileDate             time.Time `json:"fileDate"`
			FileLength           int       `json:"fileLength"`
			DownloadCount        int       `json:"downloadCount"`
			DownloadURL          string    `json:"downloadUrl"`
			GameVersions         []string  `json:"gameVersions"`
			SortableGameVersions []struct {
				GameVersionName        string    `json:"gameVersionName"`
				GameVersionPadded      string    `json:"gameVersionPadded"`
				GameVersion            string    `json:"gameVersion"`
				GameVersionReleaseDate time.Time `json:"gameVersionReleaseDate"`
				GameVersionTypeID      int       `json:"gameVersionTypeId"`
			} `json:"sortableGameVersions"`
			Dependencies []struct {
				ModID        int `json:"modId"`
				FileID       int `json:"fileId"`
				RelationType int `json:"relationType"`
			} `json:"dependencies"`
			ExposeAsAlternative bool `json:"exposeAsAlternative"`
			ParentProjectFileID int  `json:"parentProjectFileId"`
			AlternateFileID     int  `json:"alternateFileId"`
			IsServerPack        bool `json:"isServerPack"`
			ServerPackFileID    int  `json:"serverPackFileId"`
			FileFingerprint     int  `json:"fileFingerprint"`
			Modules             []struct {
				Name        string `json:"name"`
				Fingerprint int    `json:"fingerprint"`
			} `json:"modules"`
		} `json:"latestFiles"`
		LatestFilesIndexes []struct {
			GameVersion       string `json:"gameVersion"`
			FileID            int    `json:"fileId"`
			Filename          string `json:"filename"`
			ReleaseType       int    `json:"releaseType"`
			GameVersionTypeID int    `json:"gameVersionTypeId"`
			ModLoader         int    `json:"modLoader"`
		} `json:"latestFilesIndexes"`
		DateCreated  time.Time `json:"dateCreated"`
		DateModified time.Time `json:"dateModified"`
		DateReleased time.Time `json:"dateReleased"`
	} `json:"data"`
}