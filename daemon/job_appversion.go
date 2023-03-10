package daemon

import (
	"github.com/coreos/go-semver/semver"
)

func JobVersionCheck(dm *DataManager, api *RepoAPI) func() {
	return func() {
		// if no currentVersion data is available, 0.0.0-0 currentVersion will be used.
		currentVersion := semver.New(api.version)
		vdata := dm.GetVersionData()

		if currentVersion.LessThan(vdata.version) {
			dm.SetVersionData(vdata.version, true)
		}
		if vdata.newerVersionAvailable {
			return
		}

		// Get info from repo and convert data to strings
		var versionStrings []string
		switch api.packageType {
		case "deb":
			data, err := api.DebianFileList()
			if err != nil {
				dm.SetVersionData(vdata.version, false)
				return
			}
			versionStrings = ParseDebianVersions(data)
		default: // use this logic for RPM and unexpected cases
			data, err := api.RpmFileList()
			if err != nil {
				dm.SetVersionData(vdata.version, false)
				return
			}
			versionStrings = ParseRpmVersions(data)
		}
		// Convert currentVersion strings to a format that's easier to compare
		versions := StringsToVersions(versionStrings)
		latestVersion := GetLatestVersion(versions)

		if currentVersion.LessThan(latestVersion) {
			dm.SetVersionData(latestVersion, true)
		}
	}
}
