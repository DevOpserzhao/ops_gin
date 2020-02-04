package tag_service

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/DevOpserzhao/ops_gin/first/models"
	"github.com/DevOpserzhao/ops_gin/first/pkg/file"
	"github.com/DevOpserzhao/ops_gin/first/pkg/gredis"
	"github.com/DevOpserzhao/ops_gin/first/pkg/logging"
	"github.com/DevOpserzhao/ops_gin/first/service/cache_service"
	"github.com/EDDYCJY/go-gin-example/pkg/export"
	"github.com/tealeg/xlsx"
	"io"
	"strconv"
	"time"
)

type Tag struct {
	ID         int
	Name       string
	CreatedBy  string
	ModifiedBy string
	State      int

	PageNum  int
	PageSize int
}

func (t *Tag) ExistByName() (bool, error) {
	return models.ExistTagByName(t.Name)
}

func (t *Tag) ExistByID() (bool, error) {
	return models.ExistTagByID(t.ID)
}

func (t *Tag) Add() error {
	return models.AddTag(t.Name, t.State, t.CreatedBy)
}

func (t *Tag) Edit() error {
	data := make(map[string]interface{})
	data["modified_by"] = t.ModifiedBy
	data["name"] = t.Name
	if t.State >= 0 {
		data["state"] = t.State
	}

	return models.EditTag(t.ID, data)
}

func (t *Tag) Delete() error {
	return models.DeleteTag(t.ID)
}

func (t *Tag) Count() (int, error) {
	return models.GetTagTotal(t.getMaps())
}

func (t *Tag) GetAll() ([]models.Tag, error) {
	var (
		//tags []models.Tag
		tags, cacheTags []models.Tag
	)

	cache := cache_service.Tag{
		State: t.State,

		PageNum:  t.PageNum,
		PageSize: t.PageSize,
	}
	key := cache.GetTagsKey()
	fmt.Printf("tag_service key=%v\n\n", key)
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		fmt.Printf("redis data=%v\n\n", data)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheTags)
			fmt.Printf("cacheTags====== =%v\n\n", cacheTags)
			return cacheTags, nil
		}
	}

	//	fmt.Println("执行生成缓存数据")
	tags, err := models.GetTags(t.PageNum, t.PageSize, t.getMaps())
	if err != nil {
		return nil, err
	}
	fmt.Printf("执行生成缓存数据%v\n\n", tags)

	gredis.Set(key, tags, 3600)
	return tags, nil
}

func (t *Tag) Export() (string, error) {
	tags, err := t.GetAll()
	if err != nil {
		return "", err
	}

	xlsFile := xlsx.NewFile()
	sheet, err := xlsFile.AddSheet("标签信息")
	if err != nil {
		return "", err
	}

	titles := []string{"ID", "名称", "创建人", "创建时间", "修改人", "修改时间"}
	row := sheet.AddRow()

	var cell *xlsx.Cell
	for _, title := range titles {
		cell = row.AddCell()
		cell.Value = title
	}

	for _, v := range tags {
		values := []string{
			strconv.Itoa(v.ID),
			v.Name,
			v.CreatedBy,
			strconv.Itoa(v.CreatedOn),
			v.ModifiedBy,
			strconv.Itoa(v.ModifiedOn),
		}

		row = sheet.AddRow()
		for _, value := range values {
			cell = row.AddCell()
			cell.Value = value
		}
	}

	time := strconv.Itoa(int(time.Now().Unix()))
	filename := "tags-" + time + export.EXT

	dirFullPath := export.GetExcelFullPath()
	err = file.IsNotExistMkDir(dirFullPath)
	if err != nil {
		return "", err
	}

	err = xlsFile.Save(dirFullPath + filename)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func (t *Tag) Import(r io.Reader) error {
	xlsx, err := excelize.OpenReader(r)
	if err != nil {
		return err
	}

	rows := xlsx.GetRows("标签信息")
	for irow, row := range rows {
		if irow > 0 {
			var data []string
			for _, cell := range row {
				data = append(data, cell)
			}

			models.AddTag(data[1], 1, data[2])
		}
	}

	return nil
}

func (t *Tag) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0

	if t.Name != "" {
		maps["name"] = t.Name
	}
	if t.State >= 0 {
		maps["state"] = t.State
	}
	fmt.Printf("tag_service maps=%v\n\n", maps)
	return maps
}
