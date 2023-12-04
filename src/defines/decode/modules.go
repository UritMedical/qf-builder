/**
 * @Author: Joey
 * @Description:
 * @Create Date: 2023/12/4 14:53
 */

package decode

import (
	"qf-builder/defines/decode/util"
	"qf-builder/defines/fileModels"
	"qf-builder/models"
	"qf-builder/util/strHelper"
)

func FileToModules(modules []fileModels.FileModule) ([]models.Module, error) {
	var modulesList []models.Module
	for _, v := range modules {
		node, e := FileToModule(v)
		if e != nil {
			return nil, e
		}
		modulesList = append(modulesList, node)
	}
	return modulesList, nil
}
func FileToModule(file fileModels.FileModule) (module models.Module, err error) {
	module.Name, module.Comment, err = util.DecStdDefine(file.Module)
	module.Route = util.Trim(file.Route)
	if err != nil {
		return
	}
	module.Groups, err = FileToGroups(file.Groups)
	if err != nil {
		return
	}
	module.Notices, err = FileToNotices(file.Notices)
	return
}

// FileToNotices 从FileNotices转换为Notices
func FileToNotices(notices []fileModels.FileNotice) ([]models.Notice, error) {
	var noticesList []models.Notice
	for _, notice := range notices {
		node, e := fileToNotice(notice)
		if e != nil {
			return nil, e
		}
		noticesList = append(noticesList, node)
	}
	return noticesList, nil
}

func fileToNotice(fileNotice fileModels.FileNotice) (notice models.Notice, err error) {
	notice.Topic, notice.Comment, err = util.DecStdDefine(fileNotice.Notice)
	if err != nil {
		return
	}
	for _, v := range fileNotice.Params {
		param, e := util.DecProp(v)
		if e != nil {
			err = e
			return
		}
		notice.Params = append(notice.Params, param)
	}
	return
}

// FileToGroups 从FileGroups转换为Groups
func FileToGroups(list []fileModels.FileGroup) ([]models.Group, error) {
	var groupsList []models.Group
	for _, v := range list {
		node, e := fileToGroup(v)
		if e != nil {
			return nil, e
		}
		groupsList = append(groupsList, node)
	}
	return groupsList, nil
}

func fileToGroup(group fileModels.FileGroup) (output models.Group, err error) {
	output.Name, output.Comment, err = util.DecStdDefine(group.Group)
	var heads []models.Prop
	for _, h := range group.Heads {
		head, e := util.DecProp(h)
		if e != nil {
			return output, e
		}
		heads = append(heads, head)
	}
	for _, api := range group.Apis {
		var a models.Api
		a.Heads = append(a.Heads, heads...)
		a.Name, a.Route, a.ReqType, a.Comment, err = decApiDefine(api.Api)
		if err != nil {
			return
		}
		for _, h := range api.Head {
			head, e := util.DecParam(h)
			if e != nil {
				err = e
				return
			}
			a.Heads = append(a.Heads, head)
		}
		for _, p := range api.Query {
			param, e := util.DecParam(p)
			if e != nil {
				err = e
				return
			}
			a.Query = append(a.Query, param)
		}
		for _, j := range api.Json {
			param, e := util.DecParam(j)
			if e != nil {
				err = e
				return
			}
			a.Json = append(a.Json, param)
		}
		a.Return, err = decReturn(api.Return)
		if err != nil {
			return
		}
		output.Apis = append(output.Apis, a)

	}
	return output, nil
}

// api返回解码
func decReturn(r string) (p models.Prop, err error) {
	array := util.Split(r)
	if len(array) < 1 {
		return
	}
	var output models.Prop
	if len(array) > 1 {
		output.Comment = util.Trim(array[1])
	}
	tName := array[0]
	output.IsArray, output.IsPoint, output.TypeName = util.CheckType(tName)
	return output, nil
}

func decApiDefine(define string) (name, route, reqType, comment string, err error) {
	array := util.Split(define)
	if len(array) < 3 {
		return
	}
	name = strHelper.ToCamel(array[0])
	route = util.Trim(array[1])
	reqType = strHelper.ToCamel(array[2])
	if len(array) > 3 {
		comment = util.Trim(array[3])
	}
	return
}
