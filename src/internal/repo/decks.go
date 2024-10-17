package repo

import (
	"dc_haur/src/internal/model"
	"github.com/jinzhu/gorm"
)

type Decks struct {
	db *gorm.DB
}

func NewDecksRepo(db *gorm.DB) *Decks {
	return &Decks{db: db}
}

func (r *Decks) GetDecks() ([]model.Deck, error) {
	var decks []model.Deck
	if err := r.db.Find(&decks).Error; err != nil {
		return nil, err
	}
	return decks, nil
}

func (r *Decks) GetPublicDecks() ([]model.Deck, error) {
	var decks []model.Deck
	if err := r.db.Where("not hidden").Find(&decks).Error; err != nil {
		return nil, err
	}
	return decks, nil
}

func (r *Decks) GetDeckByName(name string) (*model.Deck, error) {
	var deck model.Deck
	if err := r.db.Where("name = ?", name).First(&deck).Error; err != nil {
		return nil, err
	}
	return &deck, nil
}

func (r *Decks) GetDecksByLanguage(language string) ([]model.Deck, error) {
	var decks []model.Deck
	if err := r.db.Where("language_code = ?", language).Find(&decks).Error; err != nil {
		return nil, err
	}
	return decks, nil
}

func (r *Decks) GetPublicDecksByLanguage(language string) (*[]model.Deck, error) {
	var decks []model.Deck
	if err := r.db.Where("language_code = ? AND not hidden", language).Find(&decks).Error; err != nil {
		return nil, err
	}
	return &decks, nil
}

func (r *Decks) GetDecksOpenedByPromo(clientID string) (*[]model.Deck, error) {
	var decks []model.Deck
	selectUnlocked := "SELECT deck_id from unlocked_decks WHERE client_id = ?"
	if err := r.db.Where(`hidden AND id in (`+selectUnlocked+`)`, clientID).Find(&decks).Error; err != nil {
		return nil, err
	}
	return &decks, nil
}

func (r *Decks) GetAvailableForUserDecksByLanguage(language, userID string) ([]model.Deck, error) {
	decks, err := r.GetPublicDecksByLanguage(language)
	if err != nil {
		return nil, err
	}
	unlockedDecks, err := r.GetDecksOpenedByPromo(userID)
	if err != nil {
		return nil, err
	}
	return append(*decks, *unlockedDecks...), nil
}

func (r *Decks) GetDeckByNameWithEmoji(name string) (*model.Deck, error) {
	var deck model.Deck
	if err := r.db.Where("(concat(coalesce(concat(emoji, ' '),''), name) = ?) OR name = ?", name, name).First(&deck).Error; err != nil {
		return nil, err
	}
	return &deck, nil
}

func (r *Decks) GetQuestionsCount(ID string) (int, error) {
	var res int
	err := r.db.Raw("SELECT count(*) FROM questions where level_id in (select id from levels where deck_id = ?)", ID).
		Row().Scan(&res)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (r *Decks) GetOpenedQuestionsCount(ID string, clientID string) (int, error) {
	var opened int
	sql := `SELECT count(distinct question_id) FROM questions_history 
		where client_id = ? 
		AND level_id in (select id from levels where deck_id = ?)
		AND question_id in (select id from questions)` //если вопросы будут удаляться, чтобы это учлось
	err := r.db.Raw(sql, clientID, ID).Row().Scan(&opened)
	return opened, err
}

func (r *Decks) GetHiddenDeckByPromo(promo string, clientID string) *model.Deck {
	var deck model.Deck
	sql := `SELECT * FROM decks WHERE 
		hidden 
		AND promo = ? 
		AND id NOT IN (select deck_id from unlocked_decks where client_id = ?)
		LIMIT 1`
	err := r.db.Raw(sql, promo, clientID).Scan(&deck).Error
	if err != nil {
		return nil
	}
	return &deck
}

func (r *Decks) UnlockDeck(deckID, clientID string) error {
	err := r.db.Exec("INSERT INTO unlocked_decks (client_id, deck_id) VALUES (?,?) on duplicate key update client_id=client_id",
		clientID, deckID).Error
	return err
}

func (r *Decks) TruncateUnlockedDecks() error {
	err := r.db.Exec("TRUNCATE TABLE unlocked_decks").Error
	return err
}
