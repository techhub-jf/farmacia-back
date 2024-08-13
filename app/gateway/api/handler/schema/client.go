package schema

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/techhub-jf/farmacia-back/app/domain/erring"
)

type ClientDTO struct {
	FullName      string    `json:"full_name"`
	Birth         time.Time `json:"birth"`
	Cpf           string    `json:"cpf"`
	Rg            string    `json:"rg"`
	Phone         string    `json:"phone"`
	Cep           string    `json:"cep"`
	Address       string    `json:"address"`
	AddressNumber int       `json:"address_number"`
	District      string    `json:"district"`
	City          string    `json:"city"`
	State         string    `json:"state"`
}

type ClientResponse struct {
	ID        uint   `json:"id"`
	Reference string `json:"reference"`
	FullName  string `json:"full_name"`
	Cpf       string `json:"cpf"`
	Rg        string `json:"rg"`
	Phone     string `json:"phone"`
}

type ClientQueryParams struct {
	Page     uint64
	SortBy   string
	SortType string
	Limit    uint64
}

func (cqp *ClientQueryParams) ValidateParameters(page string, sortBy string, sortType string, limit string) {
	outputPage, err := strconv.ParseUint(page, 10, 32)
	if err != nil {
		outputPage = 1
	}

	var outputSortBy string
	if sortBy == "reference" ||
		sortBy == "full_name" ||
		sortBy == "cpf" ||
		sortBy == "rg" ||
		sortBy == "phone" {
		outputSortBy = sortBy
	} else {
		outputSortBy = "id"
	}

	sortType = strings.ToUpper(sortType)

	var outputSortType string
	if sortType == "DESC" { //nolint:goconst
		outputSortType = sortType
	} else {
		outputSortType = "ASC" //nolint:goconst
	}

	outputLimit, err := strconv.ParseUint(limit, 10, 32)
	if err != nil {
		outputLimit = 10
	}

	cqp.Page = outputPage
	cqp.SortBy = outputSortBy
	cqp.SortType = outputSortType
	cqp.Limit = outputLimit
}

func (clientDTO *ClientDTO) ValidateCpf() error {
	const (
		toBeRemoved  = `[\p{P}\s]`
		matchPattern = `^[0-9]{11}$`
	)

	reg := regexp.MustCompile(toBeRemoved)

	clientDTO.Cpf = reg.ReplaceAllString(clientDTO.Cpf, "")

	reg = regexp.MustCompile(matchPattern)

	if !reg.MatchString(clientDTO.Cpf) {
		return erring.ErrClientCpfElevenDigits
	}

	const (
		firstDigitCalcNumber  = 9
		secondDigitCalcNumber = 10
	)

	if validateDigit(clientDTO.Cpf, firstDigitCalcNumber) &&
		validateDigit(clientDTO.Cpf, secondDigitCalcNumber) {
		return nil
	}

	return erring.ErrClientCpfInvalid
}

func validateDigit(cpf string, digit int) bool {
	var total int

	for i := range digit {
		multiplier := (digit + 1 - i)

		number, err := strconv.Atoi(string(cpf[i]))
		if err != nil {
			return false
		}

		total += number * multiplier
	}

	const cpfDigitNumber = 11

	r := total % cpfDigitNumber

	result := cpfDigitNumber - r
	if result >= cpfDigitNumber-1 {
		result = 0
	}

	originalDigit, err := strconv.Atoi(string(cpf[digit]))
	if err == nil && result == originalDigit {
		return true
	}

	return false
}

func (clientDTO *ClientDTO) CheckForEmptyFields() error {
	if clientDTO.FullName == "" ||
		clientDTO.Birth.IsZero() ||
		clientDTO.Cpf == "" ||
		clientDTO.Rg == "" ||
		clientDTO.Phone == "" ||
		clientDTO.Cep == "" ||
		clientDTO.Address == "" ||
		clientDTO.AddressNumber <= 0 ||
		clientDTO.District == "" ||
		clientDTO.City == "" ||
		clientDTO.State == "" {
		return erring.ErrClientEmptyFields
	}

	return nil
}

func (clientDTO *ClientDTO) ValidateID(id string) (uint, error) {
	clientID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		return 0, erring.ErrInvalidID
	}

	return uint(clientID), nil
}
