package spinrewriter

import "strings"

type Confidence string

const (
	ConfidenceLow    Confidence = "low"
	ConfidenceMedium Confidence = "medium"
	ConfidenceHigh   Confidence = "high"
)

type Option struct {
	Key   string
	Value string
}

// Your original text that you wish to rewrite or spin. This text will be analyzed by our software, its meaning will be extracted, and Spin Rewriter will rewrite it with synonyms for individual words and phrases.
func withText(text string) Option {
	return Option{
		Key:   "text",
		Value: text,
	}
}

// A list of keywords and key phrases that you do NOT want to spin
func WithProtectedTerms(terms ...string) Option {
	return Option{
		Key:   "protected_terms",
		Value: strings.Join(terms, "\n"),
	}
}

// Should Spin Rewriter automatically protect all Capitalized Words except for those in the title of your original text?
func WithAutoProtectedTerms() Option {
	return Option{
		Key:   "auto_protected_terms",
		Value: "true",
	}
}

// The confidence level of the One-Click Rewrite process.
// • low: largest number of synonyms for various words and phrases, least readable unique variations of text
// • medium: relatively reliable synonyms, usually well readable unique variations of text (default setting)
// • high: only the most reliable synonyms, perfectly readable unique variations of text
func WithConfidenceLevel(level Confidence) Option {
	return Option{
		Key:   "confidence_level",
		Value: string(level),
	}
}

// Should Spin Rewriter also spin single words inside already spun phrases?
// If set to "true", the returned spun text might contain 2 levels of nested spinning syntax.
func WithNestedSpintax() Option {
	return Option{
		Key:   "nested_spintax",
		Value: "true",
	}
}

// Should Spin Rewriter spin complete sentences?
// If set to "true", some sentences will be replaced with a (shorter) spun variation.
func WithAutoSenteces() Option {
	return Option{
		Key:   "auto_sentences",
		Value: "true",
	}
}

// Should Spin Rewriter spin entire paragraphs?
// If set to "true", some paragraphs will be replaced with a (shorter) spun variation.
func WithAutoParagraphs() Option {
	return Option{
		Key:   "auto_paragraphs",
		Value: "true",
	}
}

// Should Spin Rewriter automatically write additional paragraphs on its own?
// If set to "true", the returned spun text will contain additional paragraphs.
func WithAutoNewParagraphs() Option {
	return Option{
		Key:   "auto_new_paragraphs",
		Value: "true",
	}
}

// Should Spin Rewriter automatically change the entire structure of phrases and sentences?
// If set to "true", Spin Rewriter will change "If he is hungry, John eats." to "John eats if he is hungry." and "John eats and drinks." to "John drinks and eats."
func WithAutoSentenceTrees() Option {
	return Option{
		Key:   "auto_sentence_trees",
		Value: "true",
	}
}

// optional parameter with "unique_variation" and "unique_variation_from_spintax" API requests
// Should Spin Rewriter use only synonyms of the original words instead of the original words themselves?
// If set to "true", Spin Rewriter will never use any of the original words of phrases if there is a synonym available. This significantly improves the uniqueness of generated spun content.
func WithUseOnlySynonyms() Option {
	return Option{
		Key:   "use_only_synonyms",
		Value: "true",
	}
}

// optional parameter with "unique_variation" and "unique_variation_from_spintax" API requests
// Should Spin Rewriter intelligently randomize the order of paragraphs and unordered lists when generating spun text?
// If set to "true", Spin Rewriter will randomize the order of paragraphs and lists where possible while preserving the readability of the text. This significantly improves the uniqueness of generated spun content.
func WithReorderParagraphs() Option {
	return Option{
		Key:   "reorder_paragraphs",
		Value: "true",
	}
}

// optional parameter with "unique_variation" and "unique_variation_from_spintax" API requests
// Should Spin Rewriter humanize the text with the help of an AI LLM model?
// If set to "true", Spin Rewriter will make the generated spun content more human-like by applying an AI LLM model that has been fine-tuned on best-performing SEO content.
func WithHumanizeAI() Option {
	return Option{
		Key:   "humanize_ai",
		Value: "true",
	}
}

// optional parameter with "unique_variation" and "unique_variation_from_spintax" API requests
// Should Spin Rewriter humanize the text by adding common typos?
// If set to "true", Spin Rewriter will make the generated spun content more human-like by introducing some common typos, e.g. "teh" instead of "the".
func WithHumanizeTypos() Option {
	return Option{
		Key:   "humanize_typos",
		Value: "true",
	}
}

// optional parameter with "unique_variation" and "unique_variation_from_spintax" API requests
// Should Spin Rewriter humanize the text by adding common misspellings?
// If set to "true", Spin Rewriter will make the generated spun content more human-like by introducing some common misspellings, e.g. "seperate" instead of "separate".
func WithHumanizeMisspellings() Option {
	return Option{
		Key:   "humanize_misspellings",
		Value: "true",
	}
}

// optional parameter with "unique_variation" and "unique_variation_from_spintax" API requests
// Should Spin Rewriter humanize the text by adding common homophone confusions?
// If set to "true", Spin Rewriter will make the generated spun content more human-like by introducing some common homophone confusions, e.g. "there" instead of "their".
func WithHumanizeHomophones() Option {
	return Option{
		Key:   "humanize_homophones",
		Value: "true",
	}
}

// optional parameter with "unique_variation" and "unique_variation_from_spintax" API requests
// Should Spin Rewriter humanize the text by adding common capitalization issues?
// If set to "true", Spin Rewriter will make the generated spun content more human-like by introducing some common capitalization issues, e.g. "JOhn" instead of "John".
func WithHumanizeCapitalization() Option {
	return Option{
		Key:   "humanize_capitalization",
		Value: "true",
	}
}

// optional parameter with "unique_variation" and "unique_variation_from_spintax" API requests
// Should Spin Rewriter humanize the text by adding common punctuation issues?
// If set to "true", Spin Rewriter will make the generated spun content more human-like by introducing some common punctuation issues, e.g. "Ive" instead of "I've".
func WithHumanizePunctuation() Option {
	return Option{
		Key:   "humanize_punctuation",
		Value: "true",
	}
}

// optional parameter with "unique_variation" and "unique_variation_from_spintax" API requests
// Should Spin Rewriter humanize the text by adding repeated letters and words?
// If set to "true", Spin Rewriter will make the generated spun content more human-like by introducing some commonly repeated letters and words, e.g. "that that".
func WithHumanizeRepeats() Option {
	return Option{
		Key:   "humanize_repeats",
		Value: "true",
	}
}

// optional parameter with "unique_variation" and "unique_variation_from_spintax" API requests
// Should Spin Rewriter humanize the text by adding common spacing issues?
// If set to "true", Spin Rewriter will make the generated spun content more human-like by introducing some common spacing issues, e.g. two spaces after a comma.
func WithHumanizeSpacing() Option {
	return Option{
		Key:   "humanize_spacing",
		Value: "true",
	}
}

// optional parameter with "unique_variation" and "unique_variation_from_spintax" API requests
// Should Spin Rewriter automatically enrich generated articles with HTML tags for headings, lists with bullet points, etc.?
// If set to "true", Spin Rewriter will add HTML code for headings (<h1>, <h2>) and lists (<ul>, <li>) that it detects in generated articles. This improves the readability of generated spun content.
func WithHTMLMarkup() Option {
	return Option{
		Key:   "add_html_markup",
		Value: "true",
	}
}

// optional parameter with "unique_variation" and "unique_variation_from_spintax" API requests
// Should Spin Rewriter automatically convert line-breaks to valid HTML tags?
// If set to "true", Spin Rewriter will convert all line-breaks (newline characters) in generated unique articles into valid HTML tags (<br>). This lets you publish HTML code of generated unique articles without any additional work required.
func WithHTMLLineBreaks() Option {
	return Option{
		Key:   "use_html_linebreaks",
		Value: "true",
	}
}

// optional parameter with "text_with_spintax" API requests.
// The spintax format of the returned spun text.
// • {|}: the {first option|second option} spintax used (default setting)
// • {~}: the {first option~second option} spintax used
// • [|]: the [first option|second option] spintax used
// • [spin]: the [spin]first option|second option[/spin] spintax used
// • #SPIN: the {#SPIN: first option || second option #} spintax used
func WithSpintaxFormat(format string) Option {
	return Option{
		Key:   "spintax_format",
		Value: format,
	}
}

func optionValue(options []Option, key string, def string) string {
	for _, opt := range options {
		if opt.Key == key {
			return opt.Value
		}
	}

	return def
}
