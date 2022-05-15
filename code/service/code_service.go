package service

import (
	"context"
	"math/rand"
	"time"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain"
)

type codeService struct {
	codeRepo domain.CodeRepository
}

func NewCodeService(cr domain.CodeRepository) *codeService {
	return &codeService{codeRepo: cr}
}

func (cs *codeService) GenerateNewCode(ctx context.Context, identifier string, length int8, number bool, symbol bool) (*domain.Code, error) {
	rand.Seed(time.Now().UnixNano())
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numberRunes := []rune("1234567890")
	symbolRunes := []rune(":?=-()/%@!")

	code := &domain.Code{Identifier: identifier}

	b := make([]rune, length)

	if number && symbol {
		for i := range b {
			r := rand.Intn(3)
			if r == 0 {
				b[i] = letterRunes[rand.Intn(len(letterRunes))]
			} else if r == 1 {
				b[i] = numberRunes[rand.Intn(len(numberRunes))]
			} else {
				b[i] = symbolRunes[rand.Intn(len(symbolRunes))]
			}
		}

		code.Value = string(b)
	}

	if number && !symbol {
		for i := range b {
			r := rand.Intn(2)
			if r == 0 {
				b[i] = letterRunes[rand.Intn(len(letterRunes))]
			} else if r == 1 {
				b[i] = numberRunes[rand.Intn(len(numberRunes))]
			}
		}

		code.Value = string(b)
	}

	if !number && symbol {
		for i := range b {
			r := rand.Intn(2)
			if r == 0 {
				b[i] = letterRunes[rand.Intn(len(letterRunes))]
			} else if r == 1 {
				b[i] = symbolRunes[rand.Intn(len(symbolRunes))]
			}
		}

		code.Value = string(b)
	}

	if !number && !symbol {
		for i := range b {
			b[i] = letterRunes[rand.Intn(len(letterRunes))]
		}

		code.Value = string(b)
	}

	if err := cs.codeRepo.Store(ctx, code); err != nil {
		return nil, err
	}

	return code, nil
}

func (cs *codeService) GenerateNewCodeFake(ctx context.Context) {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration((8 + rand.Intn(5))) * time.Second)
}

func (cs *codeService) ValidateCode(ctx context.Context, c *domain.Code) (domain.IsValid, error) {
	code, err := cs.codeRepo.GetByValue(ctx, c.Value)

	if err != nil {
		return false, err
	}

	if code.Identifier == c.Identifier && code.Value == c.Value {
		if err := cs.codeRepo.DeleteByValue(ctx, c.Value); err != nil {
			return false, err
		} else {
			return true, nil
		}
	} else {
		return false, nil
	}
}
