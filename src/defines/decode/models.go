package decode

import (
	"qf-builder/defines/decode/util"
	"qf-builder/defines/fileModels"
	"qf-builder/models"
)

// FileToModels 从FileModels转换为Models
func FileToModels(list []fileModels.FileModel) ([]models.Model, error) {
	var modelsList []models.Model
	for _, v := range list {
		node, e := fileToModel(v)
		if e != nil {
			return nil, e
		}
		modelsList = append(modelsList, node)
	}
	return modelsList, nil
}

// fileToModel 将FileModel转换为Model
func fileToModel(ft fileModels.FileModel) (output models.Model, err error) {
	output = models.Model{}
	err = nil

	name, comment, e := util.DecStdDefine(ft.Model)
	if e != nil {
		err = e
		return
	}
	output.Name = name
	output.Comment = comment
	for _, v := range ft.Props {
		prop, e := util.DecProp(v)
		if e != nil {
			err = e
			return
		}
		output.Props = append(output.Props, prop)
	}
	return output, err
}
