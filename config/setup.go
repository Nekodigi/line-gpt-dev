package config

import (
	"log"

	"github.com/Nekodigi/line-gpt-dev/infrastructure"
	"github.com/Nekodigi/line-gpt-dev/models"
)

var (
	fs *infrastructure.Firestore
)

func init() {
	conf := Load()

	var err error
	fs, err = infrastructure.NewFirestore(conf.ProjectId)
	if err != nil {
		log.Fatalf("firestore.New: %+v", err)
	}
}

func Setup() {
	// cmd, err := fs.CreateCommand("会話")
	// if err != nil {
	// 	log.Fatalf("Create command: %+v", err)
	// }
	presetCmds := []*models.Command{}
	presetCmds = append(presetCmds, &models.Command{Id: "会話", Args: []string{}, Messages: `[{"Role": "user", "Content": "%s"}]`})
	presetCmds = append(presetCmds, &models.Command{Id: "励まし", Args: []string{}, Messages: `[{"Role": "user", "Content": "次の文字列に返信する形で励ましのメッセージを返信してください:%s"}]`})
	presetCmds = append(presetCmds, &models.Command{Id: "要約", Args: []string{}, Messages: `[{"Role": "user", "Content": "要約:%s"}]`})
	presetCmds = append(presetCmds, &models.Command{Id: "翻訳", Args: []string{"言語"}, Messages: `[{"Role": "user", "Content": "Translate following sentence into %s:%s"}]`})

	for _, cmd := range presetCmds {
		err := fs.SetCommand(cmd)
		if err != nil {
			log.Fatalf("Create command: %+v", err)
		}
	}
}
