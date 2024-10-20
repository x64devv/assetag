package controllers

import (
	"github.com/x64devv/assetag/database"
	"github.com/x64devv/assetag/models"
)

func AddAsset(asset models.Asset) error {
	result := database.DbInstanse.Create(&asset)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateAssetStatus(asset models.Asset) error {
	result := database.DbInstanse.Save(&asset)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAllAssets() ([]models.Asset, error) {
	var assets []models.Asset
	result := database.DbInstanse.Find(&assets)
	if result.Error != nil {
		return assets, result.Error
	}
	return assets, nil
}

func SearchAsset(searchTerm string) ([]models.Asset, error) {
	var assets []models.Asset
	searchTerm = "%" + searchTerm + "%"
	result := database.DbInstanse.Where("name like ? or serial_number like ? or asset_code like ? ", searchTerm, searchTerm, searchTerm).Find(&assets)
	if result != nil {
		return assets, result.Error
	}
	return assets, nil
}
