package main

import (
	"fmt"
	"html/template"

	uos "github.com/mschoebel/urban-octo-succotash"
)

type dataTable struct{}

func (t dataTable) Name() string {
	return "data"
}

func (t dataTable) ModelName() string {
	return "person"
}

func (t dataTable) LoadData(c uos.TableConfiguration) (uos.TableData, error) {
	var dataList []person

	err := c.LoadTable(&dataList)
	if err != nil {
		return nil, err
	}

	return uos.DBExtract(t.ModelName(), dataList, c.Columns), nil
}

func (t dataTable) ColumnInfo(columns []string) []uos.TableColumn {
	result := make([]uos.TableColumn, len(columns))

	for i, c := range columns {
		result[i] = uos.TableColumn{
			DisplayName: map[string]string{
				"id":   "ID",
				"name": "Name",
				"city": "City",
			}[c],
			IsSortable: map[string]bool{
				"id":   true,
				"name": true,
			}[c],
			Format: map[string]uos.TableFormatFunc{
				"name": func(id uint, raw interface{}) interface{} {
					if s, ok := raw.(string); ok {
						return template.HTML(
							fmt.Sprintf(
								`<a href="#" hx-get="/dialogs/data?id=%d" hx-target="body" hx-swap="beforeend">%s</a>`,
								id,
								template.HTMLEscapeString(s),
							),
						)
					}
					return raw
				},
			}[c],
		}
	}

	return result
}

func (t dataTable) ColumnDefault() []string {
	return []string{"name", "city"}
}

func (t dataTable) Actions() *uos.TableActions {
	return &uos.TableActions{
		uos.TableActionButton("plus", "Add", "").Dialog("data"),
		uos.TableActionButton("skull-crossbones", "Delete", "").Post("/tables/data?action=delete", ".selection"),
	}
}

func (t dataTable) DisplaySettings() uos.TableDisplayProperties {
	return uos.TableDisplayProperties{
		IsFullWidth:   true,
		IsHoverable:   true,
		IsStriped:     true,
		IsMobileReady: true,
		IsSelectable:  true,
	}
}

func (t dataTable) Delete(ids []uint) (*uos.ResponseAction, error) {
	return uos.ResponseRefresh(), uos.DB.Where("id IN (?)", ids).Delete(&person{}).Error
}
