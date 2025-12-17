package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
)

const (
	synodicMonth = 29.530588853 // Giorni del ciclo lunare sinodico
	knownNewMoon = 2451549.5    // Julian Day di una luna nuova conosciuta (6 Gen 2000)
)

// MoonPhase rappresenta una fase lunare
type MoonPhase struct {
	Age         float64 // Età della luna in giorni
	Illuminated float64 // Percentuale di illuminazione (0-100)
	PhaseName   string  // Nome della fase
	PhaseEmoji  string  // Emoji rappresentativa
}

// julianDay calcola il Julian Day Number per una data
func julianDay(t time.Time) float64 {
	year := t.Year()
	month := int(t.Month())
	day := t.Day()
	hour := t.Hour()
	minute := t.Minute()
	second := t.Second()

	// Converti ore, minuti, secondi in frazione di giorno
	dayFraction := float64(hour)/24.0 + float64(minute)/(24.0*60.0) + float64(second)/(24.0*3600.0)

	// Algoritmo Julian Day
	if month <= 2 {
		year--
		month += 12
	}

	a := year / 100
	b := 2 - a + a/4

	jd := math.Floor(365.25*float64(year+4716)) +
		math.Floor(30.6001*float64(month+1)) +
		float64(day) + dayFraction + float64(b) - 1524.5

	return jd
}

// calculateMoonPhase calcola la fase lunare corrente
func calculateMoonPhase(t time.Time) MoonPhase {
	jd := julianDay(t)

	// Calcola i giorni dalla luna nuova di riferimento
	daysSinceNew := jd - knownNewMoon

	// Calcola l'età della luna nel ciclo corrente
	age := math.Mod(daysSinceNew, synodicMonth)
	if age < 0 {
		age += synodicMonth
	}

	// Calcola la percentuale di illuminazione
	// Formula semplificata basata sull'età
	illuminated := (1 - math.Cos(2*math.Pi*age/synodicMonth)) / 2 * 100

	// Determina il nome della fase
	phaseName, emoji := getPhaseName(age)

	return MoonPhase{
		Age:         age,
		Illuminated: illuminated,
		PhaseName:   phaseName,
		PhaseEmoji:  emoji,
	}
}

// getPhaseName determina il nome della fase in base all'età
func getPhaseName(age float64) (string, string) {
	const (
		newMoon        = 1.84566
		firstQuarter   = 7.38264
		fullMoon       = 14.76529
		lastQuarter    = 22.14794
		waxingCrescent = 5.53
		waxingGibbous  = 12.91
		waningGibbous  = 20.30
		waningCrescent = 27.69
	)

	switch {
	case age < newMoon:
		return "Luna Nuova", "🌑"
	case age < waxingCrescent:
		return "Crescente", "🌒"
	case age < firstQuarter:
		return "Primo Quarto", "🌓"
	case age < waxingGibbous:
		return "Gibbosa Crescente", "🌔"
	case age < fullMoon:
		return "Luna Piena", "🌕"
	case age < waningGibbous:
		return "Gibbosa Calante", "🌖"
	case age < lastQuarter:
		return "Ultimo Quarto", "🌗"
	case age < waningCrescent:
		return "Calante", "🌘"
	default:
		return "Luna Nuova", "🌑"
	}
}

// getWiccaMeaning restituisce il significato Wicca della fase lunare
func getWiccaMeaning(age float64) string {
	const (
		newMoon        = 1.84566
		firstQuarter   = 7.38264
		fullMoon       = 14.76529
		lastQuarter    = 22.14794
		waxingCrescent = 5.53
		waxingGibbous  = 12.91
		waningGibbous  = 20.30
		waningCrescent = 27.69
	)

	switch {
	case age < newMoon:
		return "✨ Nuovi inizi - Meditazione e pianificazione"
	case age < waxingCrescent:
		return "✨ Crescita - Magia di attrazione e prosperità"
	case age < firstQuarter:
		return "✨ Azione - Superare ostacoli e decidere"
	case age < waxingGibbous:
		return "✨ Perfezionamento - Affinare e preparare"
	case age < fullMoon:
		return "✨ Culmine - Massimo potere magico"
	case age < waningGibbous:
		return "✨ Gratitudine - Condivisione e ringraziamenti"
	case age < lastQuarter:
		return "✨ Rilascio - Banishment e liberazione"
	case age < waningCrescent:
		return "✨ Purificazione - Chiusura cicli"
	default:
		return "✨ Nuovi inizi - Meditazione e pianificazione"
	}
}

// drawMoon disegna la luna con Cairo
func drawMoon(da *gtk.DrawingArea, cr *cairo.Context, phase MoonPhase) {
	allocation := da.GetAllocatedWidth()
	width := float64(allocation)
	height := float64(da.GetAllocatedHeight())

	centerX := width / 2
	centerY := height / 2
	radius := math.Min(width, height) / 2.5

	// Sfondo
	cr.SetSourceRGB(0.1, 0.1, 0.15)
	cr.Paint()

	// Cerchio esterno (contorno luna)
	cr.Arc(centerX, centerY, radius, 0, 2*math.Pi)
	cr.SetSourceRGB(0.9, 0.9, 0.85)
	cr.FillPreserve()
	cr.SetSourceRGB(0.7, 0.7, 0.65)
	cr.SetLineWidth(2)
	cr.Stroke()

	// Disegna l'ombra in base alla fase
	age := phase.Age
	illuminationAngle := (age / synodicMonth) * 2 * math.Pi

	// Determina quale parte è in ombra
	if age < synodicMonth/2 {
		// Luna crescente (da nuova a piena)
		// L'ombra è a sinistra
		offset := math.Cos(illuminationAngle) * radius

		cr.Save()
		cr.Rectangle(centerX-radius, centerY-radius, radius+offset, radius*2)
		cr.Clip()
		cr.Arc(centerX, centerY, radius, 0, 2*math.Pi)
		cr.SetSourceRGB(0.2, 0.2, 0.25)
		cr.Fill()
		cr.Restore()
	} else {
		// Luna calante (da piena a nuova)
		// L'ombra è a destra
		offset := math.Cos(illuminationAngle) * radius

		cr.Save()
		cr.Rectangle(centerX+offset, centerY-radius, radius-offset, radius*2)
		cr.Clip()
		cr.Arc(centerX, centerY, radius, 0, 2*math.Pi)
		cr.SetSourceRGB(0.2, 0.2, 0.25)
		cr.Fill()
		cr.Restore()
	}
}

func main() {
	gtk.Init(nil)

	// Calcola la fase lunare corrente
	now := time.Now()
	phase := calculateMoonPhase(now)

	// Crea finestra principale
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Errore creazione finestra:", err)
	}
	win.SetTitle("Fasi Lunari")
	win.SetDefaultSize(400, 550)
	win.SetResizable(false)
	win.Connect("destroy", gtk.MainQuit)

	// Box verticale principale
	vbox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	if err != nil {
		log.Fatal("Errore creazione vbox:", err)
	}
	vbox.SetMarginTop(20)
	vbox.SetMarginBottom(20)
	vbox.SetMarginStart(20)
	vbox.SetMarginEnd(20)

	// Titolo
	titleLabel, err := gtk.LabelNew("")
	if err != nil {
		log.Fatal("Errore creazione label:", err)
	}
	titleLabel.SetMarkup("<span size='x-large' weight='bold'>🌙 Fasi Lunari</span>")
	vbox.PackStart(titleLabel, false, false, 10)

	// DrawingArea per la luna
	drawingArea, err := gtk.DrawingAreaNew()
	if err != nil {
		log.Fatal("Errore creazione drawing area:", err)
	}
	drawingArea.SetSizeRequest(300, 300)
	drawingArea.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
		drawMoon(da, cr, phase)
	})
	vbox.PackStart(drawingArea, true, true, 10)

	// Info fase corrente
	phaseLabel, err := gtk.LabelNew("")
	if err != nil {
		log.Fatal("Errore creazione phase label:", err)
	}
	phaseLabel.SetMarkup(fmt.Sprintf(
		"<span size='large' weight='bold'>%s %s</span>",
		phase.PhaseEmoji,
		phase.PhaseName,
	))
	vbox.PackStart(phaseLabel, false, false, 5)

	// Dettagli
	detailsLabel, err := gtk.LabelNew("")
	if err != nil {
		log.Fatal("Errore creazione details label:", err)
	}
	detailsLabel.SetMarkup(fmt.Sprintf(
		"<span size='medium'>Illuminazione: <b>%.1f%%</b>\nEtà lunare: <b>%.1f giorni</b>\nData: <b>%s</b></span>",
		phase.Illuminated,
		phase.Age,
		now.Format("02/01/2006 15:04"),
	))
	detailsLabel.SetJustify(gtk.JUSTIFY_CENTER)
	vbox.PackStart(detailsLabel, false, false, 5)

	// Corrispondenza Wicca
	wiccaLabel, err := gtk.LabelNew("")
	if err != nil {
		log.Fatal("Errore creazione wicca label:", err)
	}
	wiccaLabel.SetMarkup(fmt.Sprintf(
		"<span size='medium' style='italic'>%s</span>",
		getWiccaMeaning(phase.Age),
	))
	wiccaLabel.SetJustify(gtk.JUSTIFY_CENTER)
	vbox.PackStart(wiccaLabel, false, false, 5)

	// Separator
	sep, err := gtk.SeparatorNew(gtk.ORIENTATION_HORIZONTAL)
	if err != nil {
		log.Fatal("Errore creazione separator:", err)
	}
	vbox.PackStart(sep, false, false, 10)

	// Bottone aggiorna
	btnBox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
	if err != nil {
		log.Fatal("Errore creazione btnbox:", err)
	}

	refreshBtn, err := gtk.ButtonNewWithLabel("🔄 Aggiorna")
	if err != nil {
		log.Fatal("Errore creazione button:", err)
	}
	refreshBtn.Connect("clicked", func() {
		now := time.Now()
		phase = calculateMoonPhase(now)

		phaseLabel.SetMarkup(fmt.Sprintf(
			"<span size='large' weight='bold'>%s %s</span>",
			phase.PhaseEmoji,
			phase.PhaseName,
		))

		detailsLabel.SetMarkup(fmt.Sprintf(
			"<span size='medium'>Illuminazione: <b>%.1f%%</b>\nEtà lunare: <b>%.1f giorni</b>\nData: <b>%s</b></span>",
			phase.Illuminated,
			phase.Age,
			now.Format("02/01/2006 15:04"),
		))

		wiccaLabel.SetMarkup(fmt.Sprintf(
			"<span size='medium' style='italic'>%s</span>",
			getWiccaMeaning(phase.Age),
		))

		drawingArea.QueueDraw()
	})
	btnBox.SetHAlign(gtk.ALIGN_CENTER)
	btnBox.PackStart(refreshBtn, false, false, 0)

	vbox.PackStart(btnBox, false, false, 0)

	win.Add(vbox)
	win.ShowAll()

	gtk.Main()
}
