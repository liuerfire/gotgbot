package ext

import (
	"github.com/PaulSonOfLars/gotgbot/types"
	"strconv"
	"encoding/json"
	"net/url"
	"github.com/pkg/errors"
	"io"
)

func (b Bot) SendStickerStr(chatId int, stickerId string) (*Message, error) {
	sticker := b.NewSendableSticker(chatId)
	sticker.FileId = stickerId
	return sticker.Send()
}

func (b Bot) SendStickerPath(chatId int, path string) (*Message, error) {
	sticker := b.NewSendableSticker(chatId)
	sticker.Path = path
	return sticker.Send()
}

func (b Bot) SendStickerReader(chatId int, reader io.Reader) (*Message, error) {
	sticker := b.NewSendableSticker(chatId)
	sticker.Reader = reader
	return sticker.Send()
}

func (b Bot) ReplyStickerStr(chatId int, stickerId string, replyToMessageId int) (*Message, error) {
	sticker := b.NewSendableSticker(chatId)
	sticker.FileId = stickerId
	sticker.ReplyToMessageId = replyToMessageId
	return sticker.Send()
}

func (b Bot) ReplyStickerPath(chatId int, path string, replyToMessageId int) (*Message, error) {
	sticker := b.NewSendableSticker(chatId)
	sticker.Path = path
	sticker.ReplyToMessageId = replyToMessageId
	return sticker.Send()
}

func (b Bot) ReplyStickerReader(chatId int, reader io.Reader, replyToMessageId int) (*Message, error) {
	sticker := b.NewSendableSticker(chatId)
	sticker.Reader = reader
	sticker.ReplyToMessageId = replyToMessageId
	return sticker.Send()
}

func (b Bot) GetStickerSet(name string) (*types.StickerSet, error) {
	v := url.Values{}
	v.Add("name", name)

	r, err := Get(b, "getStickerSet", v)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to getStickerSet")
	}
	if !r.Ok {
		return nil, errors.New(r.Description)
	}

	var ss types.StickerSet
	json.Unmarshal(r.Result, &ss)

	return &ss, nil
}

func (b Bot) UploadStickerFileStr(userId int, pngStickerId string) (*File, error) {
	uploadSticker := b.NewSendableUploadStickerFile(userId)
	uploadSticker.FileId = pngStickerId
	return uploadSticker.Send()
}

func (b Bot) UploadStickerFilePath(userId int, path string) (*File, error) {
	uploadSticker := b.NewSendableUploadStickerFile(userId)
	uploadSticker.Path = path
	return uploadSticker.Send()
}

func (b Bot) UploadStickerFileReader(userId int, reader io.Reader) (*File, error) {
	uploadSticker := b.NewSendableUploadStickerFile(userId)
	uploadSticker.Reader = reader
	return uploadSticker.Send()
}

func (b Bot) CreateNewStickerSetStr(userId int, name string, title string, pngStickerid string, emojis string) (bool, error) {
	createNew := b.NewSendableCreateNewSticker(userId, name, title, emojis)
	createNew.FileId = pngStickerid
	return createNew.Send()
}

func (b Bot) CreateNewStickerSetPath(userId int, name string, title string, path string, emojis string) (bool, error) {
	createNew := b.NewSendableCreateNewSticker(userId, name, title, emojis)
	createNew.Path = path
	return createNew.Send()
}

func (b Bot) CreateNewStickerSetReader(userId int, name string, title string, reader io.Reader, emojis string) (bool, error) {
	createNew := b.NewSendableCreateNewSticker(userId, name, title, emojis)
	createNew.Reader = reader
	return createNew.Send()
}

func (b Bot) AddStickerToSetStr(userId int, name string, pngStickerId string, emojis string) (bool, error) {
	addSticker := b.NewSendableAddStickerToSet(userId, name, emojis)
	addSticker.FileId = pngStickerId
	return addSticker.Send()
}

func (b Bot) AddStickerToSetPath(userId int, name string, path string, emojis string) (bool, error) {
	addSticker := b.NewSendableAddStickerToSet(userId, name, emojis)
	addSticker.Path = path
	return addSticker.Send()
}

func (b Bot) AddStickerToSetReader(userId int, name string, reader io.Reader, emojis string) (bool, error) {
	addSticker := b.NewSendableAddStickerToSet(userId, name, emojis)
	addSticker.Reader = reader
	return addSticker.Send()
}

func (b Bot) SetStickerPositionInSet(sticker string, position int) (bool, error) {
	v := url.Values{}
	v.Add("sticker", sticker)
	v.Add("position", strconv.Itoa(position))

	r, err := Get(b, "setStickerPositionInSet", v)
	if err != nil {
		return false, errors.Wrapf(err, "unable to setStickerPositionInSet")
	}
	if !r.Ok {
		return false, errors.New(r.Description)
	}

	var bb bool
	json.Unmarshal(r.Result, &bb)

	return bb, nil
}

func (b Bot) DeleteStickerFromSet(sticker string) (bool, error) {
	v := url.Values{}
	v.Add("sticker", sticker)

	r, err := Get(b, "deleteStickerFromSet", v)
	if err != nil {
		return false, errors.Wrapf(err, "unable to deleteStickerFromSet")
	}
	if !r.Ok {
		return false, errors.New(r.Description)
	}

	var bb bool
	json.Unmarshal(r.Result, &bb)

	return bb, nil
}