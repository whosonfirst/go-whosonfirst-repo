package repo

import (
	"errors"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-placetypes"
	"path/filepath"
	"strings"
)

type Repo interface {
}

type DataRepo struct {
	Repo
	Source    string
	Role      string
	Placetype string
	Country   string
	Region    string
	Filter    string // PLEASE DON'T CALL ME 'Filter' ...
}

func NewDataRepoFromPath(path string) (*DataRepo, error) {

	abs_path, err := filepath.Abs(path)

	if err != nil {
		return nil, err
	}

	repo := filepath.Base(abs_path)
	return NewDataRepoFromString(repo)
}

func NewDataRepoFromString(repo string) (*DataRepo, error) {

	parts := strings.Split(repo, "-")

	if len(parts) < 2 {
		return nil, errors.New("Invalid repo name (too short)")
	}

	if len(parts) > 6 {
		return nil, errors.New("Invalid repo name (too long)")
	}

	r := DataRepo{
		Source:    "",
		Role:      "",
		Placetype: "",
		Country:   "",
		Region:    "",
		Filter:    "",
	}

	r.Source = parts[0]
	r.Role = parts[1]

	if r.Role != "data" {
		return nil, errors.New("Unsupported role")
	}

	if len(parts) > 2 {

		placetype := parts[2]

		if !placetypes.IsValidPlacetype(placetype) {
			return nil, errors.New("Invalid placetype")
		}

		r.Placetype = placetype
	}

	if len(parts) > 3 {

		country := parts[3]

		if len(country) != 2 {
			return nil, errors.New("Invalid country code")
		}

		// to do: validate country code

		r.Country = country
	}

	if len(parts) > 4 {

		region := parts[4]

		if len(region) != 2 {
			return nil, errors.New("Invalid region code")
		}

		// to do: validate region code

		r.Region = region
	}

	if len(parts) > 5 {

		filter := parts[5]
		r.Filter = filter
	}

	return &r, nil
}

func (r *DataRepo) String() string {

	parts := make([]string, 0)

	parts = append(parts, r.Source)
	parts = append(parts, r.Role)

	if r.Placetype != "" {
		parts = append(parts, r.Placetype)
	}

	if r.Country != "" {
		parts = append(parts, r.Country)
	}

	if r.Region != "" {
		parts = append(parts, r.Region)
	}

	if r.Filter != "" {
		parts = append(parts, r.Filter)
	}

	return strings.Join(parts, "-")
}

func (r *DataRepo) MetaFilename(args ...string) string {

	var pt string

	if len(args) == 0 || args[0] == "" {
		if r.Placetype == "" {
			pt = "all"
		} else {
			pt = r.Placetype
		}
	} else {
		pt = args[0]
	}

	template := r.MetaFilenameTemplate()
	return fmt.Sprintf(template, pt)
}

func (r *DataRepo) MetaFilenameTemplate() string {

	parts := make([]string, 0)

	// unfortunately this is still-necessary legacy code...
	// (20170726/thisisaaronland)

	if r.Source == "whosonfirst" {
		parts = append(parts, "wof")
	} else {
		parts = append(parts, r.Source)
	}

	parts = append(parts, "%s")

	if r.Country != "" {
		parts = append(parts, r.Country)
	}

	if r.Region != "" {
		parts = append(parts, r.Region)
	}

	if r.Filter != "" {
		parts = append(parts, r.Filter)
	}

	parts = append(parts, "latest.csv")
	return strings.Join(parts, "-")
}

func (r *DataRepo) ConcordancesFilename(args ...string) string {

	var pt string

	if len(args) == 0 || args[0] == "" {
		if r.Placetype == "" {
			pt = "all"
		} else {
			pt = r.Placetype
		}
	} else {
		pt = args[0]
	}

	template := r.ConcordancesFilenameTemplate(pt)
	return fmt.Sprintf(template, pt)
}

func (r *DataRepo) ConcordancesFilenameTemplate(pt string) string {

	parts := make([]string, 0)

	// unfortunately this is still-necessary legacy code...
	// (20170726/thisisaaronland)

	if r.Source == "whosonfirst" {
		parts = append(parts, "wof")
	} else {
		parts = append(parts, r.Source)
	}

	parts = append(parts, "%s")

	if r.Country != "" {
		parts = append(parts, r.Country)
	}

	if r.Region != "" {
		parts = append(parts, r.Region)
	}

	if r.Filter != "" {
		parts = append(parts, r.Filter)
	}

	parts = append(parts, "concordances")
	parts = append(parts, "latest.csv")

	return strings.Join(parts, "-")
}
