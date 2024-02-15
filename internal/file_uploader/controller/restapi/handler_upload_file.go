package restapi

import (
	"context"
	"fmt"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/file_uploader/model/errorenum"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/file_uploader/model/vo"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework/model/payload"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework/util"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/logger"
	"github.com/google/uuid"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/a-aslani/golang_agency_clean_architecture/internal/file_uploader/usecase/upload_file"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework"
	"github.com/gin-gonic/gin"
)

// UploadFile godoc
// @Summary uploading file
// @Schemes
// @Description uploading file
// @Tags UploadFile
// @Accept       mpfd
// @Produce      json
// @Param        file formData file true "upload file"
// @Success      200  {object}  upload_file.InportResponse
// @Failure      400  {object}  payload.Response
// @Failure      500  {object}  payload.Response
// @Router       /file-uploader/v1/upload [post]
func (r *controller) uploadFileHandler() gin.HandlerFunc {

	type InportRequest = upload_file.InportRequest
	type InportResponse = upload_file.InportResponse

	inport := framework.GetInport[InportRequest, InportResponse](r.GetUsecase(InportRequest{}))

	type request struct {
		File *multipart.FileHeader `form:"file" json:"file"`
	}

	type response struct {
		*InportResponse
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		var jsonReq request
		err := c.Bind(&jsonReq)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		dir := upload_file.Dir

		if r.cfg.TestMode {
			dir = fmt.Sprintf("../../../.%s", dir)
		}

		fileName, dst, err := uploadFile(c, dir, upload_file.UploadPath, jsonReq.File, upload_file.AllowedExtensions)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req InportRequest
		req.Name = vo.FileName(fileName)
		req.Path = vo.FilePath(dst)
		req.ID = uuid.New().String()
		req.Now = time.Now()

		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		jsonRes.InportResponse = res

		r.log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}

func uploadFile(c *gin.Context, dir, path string, file *multipart.FileHeader, mimTypes []string) (string, string, error) {

	if file == nil {
		return "", "", errorenum.FilePathIsRequired
	}

	rand.Seed(time.Now().UnixNano())

	strData := time.Now().UTC().Format("20060102_150405")

	splFileName := strings.Split(file.Filename, ".")
	mim := splFileName[len(splFileName)-1]

	fileName := strData + "_" + strconv.Itoa(int(time.Now().Unix())) + "_" + strconv.Itoa(rand.Intn(100000))
	dst := path + "/" + fileName + "." + mim

	if len(mimTypes) > 0 {
		hasTrueMim := false
		for _, allowMim := range mimTypes {
			if allowMim == strings.ToLower(mim) {
				hasTrueMim = true
				break
			}
		}

		if !hasTrueMim {
			return fileName, dst, errorenum.InvalidTypeError.Var(mim)
		}
	}

	if _, err := os.Stat(dir + path); os.IsNotExist(err) {
		_ = os.MkdirAll(dir+path, 0700)
	}

	return fileName, "/file-uploader" + dst, c.SaveUploadedFile(file, dir+dst)
}
