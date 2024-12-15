package dat

import (
	"fmt"
	"os"
	"strings"

	"github.com/v2rayA/v2rayA/core/v2ray/asset"
	"github.com/v2rayA/v2rayA/pkg/util/log"
)

func UpdateLocalGeoSite() (err error) {
	pathSiteDat, err := asset.GetV2rayLocationAsset("geosite.dat")
	if err != nil {
		return err
	}

	if err = asset.Download("https://ghp.ci/https://github.com/v2fly/domain-list-community/releases/latest/download/dlc.dat", pathSiteDat+".new"); err != nil {
		log.Warn("UpdateLocalGeoSite: %v", err)
		return
	}
	siteDatSha256, err := httpGet("https://ghp.ci/https://github.com/v2fly/domain-list-community/releases/latest/download/dlc.dat.sha256sum")
	if err != nil {
		err = fmt.Errorf("%w: %v", FailCheckSha, err)
		log.Warn("UpdateLocalGeoSite: %v", err)
		return err
	}
	var sha256 string
	if fields := strings.Fields(siteDatSha256); len(fields) != 0 {
		sha256 = fields[0]
	}
	if !checkSha256(pathSiteDat+".new", sha256) {
		err = fmt.Errorf("UpdateLocalGeoSite: %v", DamagedFile)
		return
	}
	if err := os.Rename(pathSiteDat+".new", pathSiteDat); err != nil {
		return err
	}
	log.Info("download[geosite.dat]: SUCCESS\n")
	return
}
