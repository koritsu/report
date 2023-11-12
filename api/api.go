package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"io"
	"net/http"
	"restapi-go/database"
	"restapi-go/model/entities"
	"restapi-go/model/vo"
	"strconv"
	"time"
)

// Create
// @Summary 파일을 기반으로 레코드를 생성한다.
// @Description file로 binary 파일을 받고, name이 있을 경우 name을 파일명으로, 없을 경우 file의 file명을 테이블에 저장한다.
// @Accept json
// @Produce json
// @Param file formData file true "content file (100MB 초과 불가)"
// @Param name formData string false "파일명"
// @Router /create [post]
func Create(context *gin.Context) {

	formFile, err := context.FormFile("file")
	if err != nil {
		commonResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	if formFile.Size > 1024*1024*100 {
		commonResponse(context, http.StatusBadRequest, "file size is over 100MB")
		return
	}

	name, ok := context.GetPostForm("name")
	if !ok {
		name = formFile.Filename
	}

	file, err := formFile.Open()

	buff := make([]byte, 1024)
	var body []byte

	// 루프
	for {
		// 읽기
		cnt, err := file.Read(buff)
		if err != nil && err != io.EOF {
			panic(err)
		}

		// 끝이면 루프 종료
		if cnt == 0 {
			break
		}

		// 쓰기
		body = append(body, buff[:cnt]...)
	}

	db := database.DB()

	var content = entities.Content{
		Content: body,
		Name:    name,
		Created: time.Now(),
	}

	tx := db.Create(&content)

	tx.Commit()

	commonResponse(context, http.StatusOK, content.ID)

}

// Read
// @Summary path id로 데이터를 조회
// @Description 아이디, 파일명, 생성일시를 제공한다
// @Accept json
// @Produce json
// @Param id path int true "content Id"
// @Success 200 {object} vo.GetContentResponse
// @Router /read/{id} [get]
func Read(context *gin.Context) {

	strId := context.Param("id")

	id, err := strconv.Atoi(strId)

	if err != nil {
		commonResponse(context, http.StatusUnprocessableEntity, "Invalid id")
		return
	}

	db := database.DB()

	content := entities.Content{}

	tx := db.First(&content, id)

	if tx.Error != nil {
		commonResponse(context, http.StatusNotFound, tx.Error.Error())
		return
	}

	context.JSON(http.StatusOK,
		vo.GetContentResponse{
			Id:      content.ID,
			Created: content.Created,
			Name:    content.Name,
		})
}

// ReadFile
// @Summary path id로 파일 조회
// @Description id의 content를 파일로 다운로드 한다.
// @Accept json
// @Produce json
// @Param id path int true "content Id"
// @Success 200 {object} vo.GetContentResponse
// @Router /read/{id}/file [get]
func ReadFile(context *gin.Context) {

	strId := context.Param("id")

	id, err := strconv.Atoi(strId)

	if err != nil {
		commonResponse(context, http.StatusUnprocessableEntity, "Invalid id")
		return
	}

	db := database.DB()

	content := entities.Content{}

	tx := db.First(&content, id)

	if tx.Error != nil {
		commonResponse(context, http.StatusNotFound, tx.Error.Error())
		return
	}

	context.Header("Content-Description", "File Transfer")
	context.Header("Content-Transfer-Encoding", "binary")
	context.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", content.Name))
	context.Data(http.StatusOK, "application/octet-stream", content.Content)
}

// Update
// @Summary path id로 수정한다.
// @Description 파일과 name을 수정한다.
// @Accept json
// @Produce json
// @Param id path int true "content Id"
// @Param file formData file true "content file"
// @Param name formData string false "파일명"
// @Router /update/{id} [put]
func Update(context *gin.Context) {
	strId := context.Param("id")

	id, err := strconv.Atoi(strId)

	if err != nil {
		commonResponse(context, http.StatusUnprocessableEntity, "Invalid id")
		return
	}

	formFile, err := context.FormFile("file")
	if err != nil {
		commonResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	name, ok := context.GetPostForm("name")
	if !ok {
		name = formFile.Filename
	}

	file, err := formFile.Open()

	buff := make([]byte, 1024)

	var body []byte
	// 루프
	for {
		// 읽기
		cnt, err := file.Read(buff)
		if err != nil && err != io.EOF {
			panic(err)
		}

		// 끝이면 루프 종료
		if cnt == 0 {
			break
		}

		// 쓰기
		body = append(body, buff[:cnt]...)
	}

	db := database.DB()

	content := entities.Content{}

	tx := db.First(&content, id)

	if tx.Error != nil {
		commonResponse(context, http.StatusNotFound, tx.Error.Error())
		return
	}

	content.Name = name
	content.Content = body

	db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&content)

	commonResponse(context, http.StatusOK, content.ID)

}

// Delete
// @Summary path id로 레코드를 삭제한다.
// @Description 삭제한다.
// @Accept json
// @Produce json
// @Param id path int true "content Id"
// @Router /delete/{id} [delete]
func Delete(context *gin.Context) {
	strId := context.Param("id")

	id, err := strconv.Atoi(strId)

	if err != nil {
		commonResponse(context, http.StatusUnprocessableEntity, "Invalid id")
		return
	}

	db := database.DB()

	content := entities.Content{}

	tx := db.First(&content, id)

	if tx.Error != nil {
		commonResponse(context, http.StatusNotFound, tx.Error.Error())
		return
	}

	tx = db.Delete(&content)

	if tx.Error != nil {
		commonResponse(context, http.StatusNotFound, tx.Error.Error())
		return
	}
	commonResponse(context, http.StatusOK, content.ID)
}

func commonResponse(context *gin.Context, code int, message any) {
	context.JSON(code, gin.H{"message": message})
}
