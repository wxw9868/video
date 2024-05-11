package service

import (
	"errors"
	"strconv"
	"time"

	"github.com/wxw9868/video/model"
	"github.com/wxw9868/video/utils"
	"gorm.io/gorm"
)

type VideoService struct{}

type Video struct {
	ID            uint    `json:"id"`
	Title         string  `json:"title"`
	Actress       string  `json:"actress"`
	Size          float64 `json:"size"`
	Duration      string  `json:"duration"`
	ModTime       string  `json:"mod_time"`
	Poster        string  `json:"poster"`
	Width         int     `json:"width"`
	Height        int     `json:"height"`
	CodecName     string  `json:"codec_name"`
	ChannelLayout string  `json:"channel_layout"`
}

func (as *VideoService) Find(actressID string) ([]Video, error) {
	dbVideo := db.Model(&model.Video{})
	if actressID != "" {
		var actress model.Actress
		if err := db.Select("Actress").Where("id = ?", actressID).First(&actress).Error; err != nil {
			return nil, err
		}
		dbVideo = dbVideo.Where("actress = ?", actress.Actress)
	}

	rows, err := dbVideo.Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var videos []Video
	for rows.Next() {
		var modelVideo model.Video
		db.ScanRows(rows, &modelVideo)

		f, _ := strconv.ParseFloat(strconv.FormatInt(modelVideo.Size, 10), 64)
		videos = append(videos, Video{
			ID:            modelVideo.ID,
			Title:         modelVideo.Title,
			Actress:       modelVideo.Actress,
			Size:          f / 1024 / 1024,
			Duration:      utils.ResolveTime(uint32(modelVideo.Duration)),
			ModTime:       modelVideo.CreationTime.Format("2006-01-02 15:04:05"),
			Poster:        modelVideo.Poster,
			Width:         modelVideo.Width,
			Height:        modelVideo.Height,
			CodecName:     modelVideo.CodecName,
			ChannelLayout: modelVideo.ChannelLayout,
		})
	}
	return videos, nil
}

func (vs *VideoService) First(id string) (model.Video, error) {
	var video model.Video
	if err := db.Where("id = ?", id).First(&video).Error; err != nil {
		return video, err
	}
	return video, nil
}

type VideoInfo struct {
	ID            uint      `json:"id"`
	Title         string    `json:"title" gorm:"column:title;type:varchar(255);comment:标题"`
	Actress       string    `json:"actress" gorm:"column:actress;type:varchar(100);comment:演员"`
	Size          int64     `json:"size" gorm:"column:size;type:bigint;comment:大小"`
	Duration      float64   `json:"duration" gorm:"column:duration;type:float;default:0;comment:时长"`
	Poster        string    `json:"poster" gorm:"column:poster;type:varchar(255);comment:封面"`
	Width         int       `json:"width" gorm:"column:width;type:int;default:0;comment:宽"`
	Height        int       `json:"height" gorm:"column:height;type:int;default:0;comment:高"`
	CodecName     string    `json:"codec_name" gorm:"column:codec_name;type:varchar(90);comment:编解码器"`
	ChannelLayout string    `json:"channel_layout" gorm:"column:channel_layout;type:varchar(90);comment:音频声道"`
	CreationTime  time.Time `gorm:"column:creation_time;type:date;comment:时间"`
	Collect       uint      `json:"collect" gorm:"column:collect;type:uint;not null;default:0;comment:收藏"`
	Browse        uint      `json:"browse" gorm:"column:browse;type:uint;not null;default:0;comment:浏览"`
	Zan           uint      `json:"zan" gorm:"column:zan;type:uint;not null;default:0;comment:赞"`
	Cai           uint      `json:"cai" gorm:"column:cai;type:uint;not null;default:0;comment:踩"`
	Watch         uint      `json:"watch" gorm:"column:watch;type:uint;not null;default:0;comment:观看"`
}

func (vs *VideoService) Info(id uint) (VideoInfo, error) {
	var videoInfo VideoInfo
	if err := db.Table("video_Video as v").Select("*,l.collect, l.browse, l.zan, l.cai, l.watch").Joins("left join video_VideoLog l on l.video_id = v.id").Where("v.id = ?", id).Scan(&videoInfo).Error; err != nil {
		return VideoInfo{}, err
	}
	return videoInfo, nil
}

func (vs *VideoService) List() ([]model.Video, error) {
	var videos []model.Video
	if err := db.Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

func (vs *VideoService) Create(videos []model.Video) error {
	if err := db.Create(&videos).Error; err != nil {
		return err
	}
	return nil
}

func (vs *VideoService) Collect(videoID uint, collect int, userID uint) error {
	var video model.Video
	result := db.First(&video, videoID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("视频不存在！")
	}

	tx := db.Begin()

	var expr string
	if collect == 1 {
		// 增加1
		expr = "collect + 1"

		if err := tx.Create(&model.UserCollectLog{UserID: userID, VideoID: videoID}).Error; err != nil {
			tx.Rollback()
			return errors.New("创建失败！")
		}
	} else {
		// 减少1
		expr = "collect - 1"

		if err := tx.Where("user_id = ? and video_id = ?", userID, videoID).Delete(&model.UserCollectLog{}).Error; err != nil {
			tx.Rollback()
			return errors.New("删除失败！")
		}
	}
	result = tx.Model(&model.VideoLog{}).Where("video_id = ?", videoID).Update("collect", gorm.Expr(expr))
	if result.Error != nil {
		tx.Rollback()
		return errors.New("更新失败！")
	}

	tx.Commit()

	return nil
}

func (vs *VideoService) Browse(videoID uint, userID uint) error {
	var video model.Video
	result := db.First(&video, videoID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("视频不存在！")
	}

	tx := db.Begin()

	if err := tx.Where(model.UserBrowseLog{UserID: userID, VideoID: videoID}).Assign("browse = ?", gorm.Expr("browse + 1")).FirstOrCreate(&model.UserBrowseLog{}).Error; err != nil {
		tx.Rollback()
		return errors.New("创建失败！")
	}
	result = tx.Model(&model.VideoLog{}).Where("video_id = ?", videoID).Update("browse", gorm.Expr("browse + 1"))
	if result.Error != nil {
		tx.Rollback()
		return errors.New("更新失败！")
	}

	tx.Commit()

	return nil
}
