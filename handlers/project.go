package handlers

import (
	"net/http"
	"regexp"
	"strings"

	"gahmen-api/helpers"
)

// @Summary List projects by ministry ID
// @Description Get a list of projects for a specific ministry
// @Tags projects
// @Produce  json
// @Param ministry_id path int true "Ministry ID"
// @Success 200 {array} types.MinistryProject "OK" "[{"project_title": "Project Alpha", "ministry_id": 1, "project_id": 1}]"
// @Router /api/v1/projects/ministry/{ministry_id} [get]
// @Security BearerAuth
func (h *Handler) ListProjectByMinistryID(w http.ResponseWriter, r *http.Request) error {
	ministry_id, err := helpers.GetIntByResponseField(r, "ministry_id")
	if err != nil {
		return err
	}
	projects, err := h.store.ListProjectsByMinistryID(ministry_id)
	if err != nil {
		return err
	}
	return helpers.WriteJSON(w, http.StatusOK, projects)
}

// @Summary Get project expenditure by ID
// @Description Get project expenditure by ID
// @Tags projects
// @Produce  json
// @Param project_id path int true "Project ID"
// @Success 200 {object} types.ProjectExpenditure "OK" "{"project_id": 1, "project_title": "Project Alpha", "ministry": "Ministry of Finance", "value_type": "Actual", "value_amount": 500000.0, "value_year": 2023, "parent_header": "Infrastructure", "document_year": 2023, "ministry_id": 1, "document_id": 1, "expenditure_id": 1}"
// @Router /api/v1/projects/{project_id} [get]
// @Security BearerAuth
func (h *Handler) GetProjectExpenditureByID(w http.ResponseWriter, r *http.Request) error {
	project_id, err := helpers.GetIntByResponseField(r, "project_id")
	if err != nil {
		return err
	}
	documents, err := h.store.GetProjectExpenditureByID(project_id)
	if err != nil {
		return err
	}
	return helpers.WriteJSON(w, http.StatusOK, documents)
}

// @Summary Search project expenditure
// @Description Search project expenditure by query with pagination
// @Tags projects
// @Produce  json
// @Param query query string false "Search query (supports AND, OR, NOT, and quoted phrases)"
// @Param id query int false "Project ID for exact match"
// @Param limit query int false "Number of results to return (default: 10)"
// @Param offset query int false "Number of results to skip (default: 0)"
// @Success 200 {array} types.ProjectExpenditure
// @Router /api/v1/projects [post]
// @Security BearerAuth
func (h *Handler) GetProjectExpenditureByQuery(w http.ResponseWriter, r *http.Request) error {
	id, idErr := helpers.GetIntByResponseQuery(r, "id")

	if idErr == nil && id != -1 {
		// If ID is provided and successfully parsed, search by ID
		documents, err := h.store.GetProjectExpenditureByID(id)
		if err != nil {
			return err
		}
		return helpers.WriteJSON(w, http.StatusOK, documents)
	}

	// If ID was not provided or had an error, then we expect a 'query'
	query, queryErr := helpers.GetStringByResponseQuery(r, "query")
	if queryErr != nil {
		return queryErr // Return error if 'query' is also missing
	}

	limit, err := helpers.GetIntByResponseQuery(r, "limit")
	if err != nil {
		limit = 10 // Default limit
	}
	offset, err := helpers.GetIntByResponseQuery(r, "offset")
	if err != nil {
		offset = 0 // Default offset
	}

	// Regex to match words, quoted phrases, and boolean operators
	// Matches:
	//   - Quoted strings: "..."
	//   - NOT keyword: NOT
	//   - OR keyword: OR
	//   - AND keyword: AND
	//   - Any other word characters
	re := regexp.MustCompile(`"([^"]*)"|NOT|OR|AND|\S+`)
	matches := re.FindAllString(query, -1)

	var processedTerms []string
	lastWasWord := false // Flag to track if the previous term processed was a regular word or quoted phrase

	for _, term := range matches {
		upperTerm := strings.ToUpper(term)
		switch upperTerm {
		case "AND":
			processedTerms = append(processedTerms, "&")
			lastWasWord = false // Operator, so reset flag
		case "OR":
			processedTerms = append(processedTerms, "|")
			lastWasWord = false // Operator, so reset flag
		case "NOT":
			processedTerms = append(processedTerms, "!")
			lastWasWord = false // Operator, so reset flag
		default:
			if strings.HasPrefix(term, "\"") && strings.HasSuffix(term, "\"") {
				// It's a quoted phrase
				if lastWasWord {
					processedTerms = append(processedTerms, "&") // Add AND if previous was a word/phrase
				}
				processedTerms = append(processedTerms, "'"+strings.Trim(term, "\"")+"'")
				lastWasWord = true // Quoted phrase is also a "word" for implicit AND
			} else {
				// Regular word
				if lastWasWord {
					processedTerms = append(processedTerms, "&") // Add AND if previous was a word/phrase
				}
				processedTerms = append(processedTerms, term+":*")
				lastWasWord = true
			}
		}
	}

	processedQuery := strings.Join(processedTerms, " ") // Join with space, to_tsquery handles operators

	documents, err := h.store.GetProjectExpenditureByQuery(processedQuery, limit, offset)
	if err != nil {
		return err
	}
	return helpers.WriteJSON(w, http.StatusOK, documents)
}

// @Summary List project expenditure by ministry ID
// @Description List project expenditure by ministry ID
// @Tags projects
// @Produce  json
// @Param ministry_id path int true "Ministry ID"
// @Success 200 {array} types.ProjectExpenditure "OK" "[{"project_id": 1, "project_title": "Project Alpha", "ministry": "Ministry of Finance", "value_type": "Actual", "value_amount": 500000.0, "value_year": 2023, "parent_header": "Infrastructure", "document_year": 2023, "ministry_id": 1, "document_id": 1, "expenditure_id": 1}]"
// @Security BearerAuth
func (h *Handler) ListProjectExpenditureByMinistryID(w http.ResponseWriter, r *http.Request) error {
	ministry_id, err := helpers.GetIntByResponseField(r, "ministry_id")
	if err != nil {
		return err
	}
	projects, err := h.store.ListProjectExpenditureByMinistryID(ministry_id)
	if err != nil {
		return err
	}
	return helpers.WriteJSON(w, http.StatusOK, projects)
}
