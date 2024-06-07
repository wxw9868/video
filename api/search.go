package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wxw9868/util"
	"github.com/wxw9868/video/utils"
)

var client = utils.NewHttpClient("http://127.0.0.1:5678/api")

type Index struct {
	Id       uint32      `json:"id" binding:"required"`
	Text     string      `json:"text" binding:"required"`
	Document interface{} `json:"document" binding:"required"`
}

func SearchIndex(c *gin.Context) {
	var bind Index
	if err := c.ShouldBindJSON(&bind); err != nil {
		c.JSON(http.StatusBadRequest, util.Fail(err.Error()))
		return
	}

	b, err := json.Marshal(&bind)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.Fail(err.Error()))
		return
	}
	Post(c, utils.Join("/index", "?", "database=default"), bytes.NewReader(b))
}

func SearchIndexBatch(c *gin.Context) {
	var binds []Index
	if err := c.ShouldBindJSON(&binds); err != nil {
		c.JSON(http.StatusBadRequest, util.Fail(err.Error()))
		return
	}

	b, err := json.Marshal(&binds)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.Fail(err.Error()))
		return
	}
	Post(c, utils.Join("/index/batch", "?", "database=default"), bytes.NewReader(b))
}

type IndexRemove struct {
	Id uint32 `json:"id" binding:"required"`
}

func SearchIndexRemove(c *gin.Context) {
	var bind IndexRemove
	if err := c.ShouldBindJSON(&bind); err != nil {
		c.JSON(http.StatusBadRequest, util.Fail(err.Error()))
		return
	}

	b, err := json.Marshal(&bind)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.Fail(err.Error()))
		return
	}
	Post(c, utils.Join("/index/remove", "?", "database=default"), bytes.NewReader(b))
}

type Query struct {
	Query     string      `json:"query" binding:"required"`
	Page      int         `json:"page"`
	Limit     int         `json:"limit"`
	Order     string      `json:"order"`
	Highlight interface{} `json:"highlight"`
	ScoreExp  string      `json:"scoreExp"`
}

func SearchQuery(c *gin.Context) {
	var bind Query
	if err := c.ShouldBindJSON(&bind); err != nil {
		c.JSON(http.StatusBadRequest, util.Fail(err.Error()))
		return
	}

	b, err := json.Marshal(&bind)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.Fail(err.Error()))
		return
	}
	Post(c, utils.Join("/query", "?", "database=default"), bytes.NewReader(b))
}

func SearchStatus(c *gin.Context) {
	Get(c, "/status")
}

func SearchDbDrop(c *gin.Context) {
	Get(c, utils.Join("/db/drop", "?", "database=", c.Query("database")))
}

func SearchWordCut(c *gin.Context) {
	Get(c, utils.Join("/word/cut", "?", "q=", c.Query("q")))
}

func Post(c *gin.Context, url string, body io.Reader) {
	resp, err := client.POST(url, "application/json", body)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.Fail(err.Error()))
		return
	}
	defer resp.Body.Close()

	robots, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.Success("SUCCESS", string(robots)))
}

func Get(c *gin.Context, url string) {
	resp, err := client.GET(url)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.Fail(err.Error()))
		return
	}
	defer resp.Body.Close()

	robots, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.Success("SUCCESS", string(robots)))
}
