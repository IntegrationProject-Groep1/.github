// badge-generator generates flat SVG badge files for the ShiftFestival org profile.
// Output is written to ../profile/badges/.
// Run: go run main.go
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// charWidth returns the approximate pixel width of a rune in DejaVu Sans 11px.
func charWidth(r rune) float64 {
	switch r {
	case 'f', 'i', 'j', 'l', 'r', 't', ' ':
		return 5.5
	case 'm', 'w', 'W', 'M':
		return 9.5
	case 'I', '|', '!', '.', ',', ';', ':':
		return 4.0
	default:
		return 7.0
	}
}

// textWidth returns the approximate rendered width for a string.
func textWidth(s string) float64 {
	w := 0.0
	for _, r := range s {
		w += charWidth(r)
	}
	return w
}

// badge produces a shields.io-style flat badge SVG.
func badge(label, message, labelColor, msgColor string) string {
	const padding = 10.0
	lw := textWidth(label) + padding
	mw := textWidth(message) + padding
	total := lw + mw
	lx := lw / 2
	mx := lw + mw/2

	return fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="%.0f" height="20" role="img" aria-label="%s: %s">
  <title>%s: %s</title>
  <linearGradient id="s" x2="0" y2="100%%">
    <stop offset="0" stop-color="#bbb" stop-opacity=".1"/>
    <stop offset="1" stop-opacity=".1"/>
  </linearGradient>
  <clipPath id="r">
    <rect width="%.0f" height="20" rx="3" fill="#fff"/>
  </clipPath>
  <g clip-path="url(#r)">
    <rect width="%.0f" height="20" fill="%s"/>
    <rect x="%.0f" width="%.0f" height="20" fill="%s"/>
    <rect width="%.0f" height="20" fill="url(#s)"/>
  </g>
  <g fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="11">
    <text aria-hidden="true" x="%.1f" y="15" fill="#010101" fill-opacity=".3">%s</text>
    <text x="%.1f" y="14">%s</text>
    <text aria-hidden="true" x="%.1f" y="15" fill="#010101" fill-opacity=".3">%s</text>
    <text x="%.1f" y="14">%s</text>
  </g>
</svg>`,
		total, label, message,
		label, message,
		total,
		total, labelColor,
		lw, mw, msgColor,
		total,
		lx+0.5, label, lx, label,
		mx+0.5, message, mx, message,
	)
}

func writeFile(path, content string) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		panic(err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		panic(err)
	}
	fmt.Printf("✓ %s\n", path)
}

func main() {
	out := "../profile/badges"

	type badgeDef struct {
		file, label, message, labelColor, msgColor string
	}

	badges := []badgeDef{
		{"project.svg", "ShiftFestival", "Integration Platform", "#2D3748", "#6B46C1"},
		{"frontend.svg", "Frontend", "Drupal · :30020", "#2D3748", "#0678BE"},
		{"kassa.svg", "Kassa", "Odoo · :8069", "#2D3748", "#714B67"},
		{"facturatie.svg", "Facturatie", "FOSSBilling · :80", "#2D3748", "#E63946"},
		{"crm.svg", "CRM", "Salesforce · :3000", "#2D3748", "#00A1E0"},
		{"planning.svg", "Planning", "Python · :30050", "#2D3748", "#3776AB"},
		{"monitoring.svg", "Monitoring", "ELK Stack · :30061", "#2D3748", "#00BFB3"},
		{"heartbeat.svg", "Heartbeat", "Python Sidecar", "#2D3748", "#38A169"},
		{"rabbitmq.svg", "Messaging", "RabbitMQ · :30000", "#2D3748", "#FF6600"},
	}

	for _, b := range badges {
		svg := badge(b.label, b.message, b.labelColor, b.msgColor)
		writeFile(filepath.Join(out, b.file), svg)
	}

	fmt.Println("All badges generated.")
}
