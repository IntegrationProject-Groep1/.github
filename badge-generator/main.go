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

// banner produces the full-width welcome banner SVG for the org profile header.
func banner() string {
	return `<svg xmlns="http://www.w3.org/2000/svg" width="960" height="160" viewBox="0 0 960 160">
  <defs>
    <linearGradient id="bg" x1="0%" y1="0%" x2="100%" y2="100%">
      <stop offset="0%" style="stop-color:#0f172a"/>
      <stop offset="100%" style="stop-color:#1e293b"/>
    </linearGradient>
    <linearGradient id="accent" x1="0%" y1="0%" x2="0%" y2="100%">
      <stop offset="0%" style="stop-color:#6366f1"/>
      <stop offset="100%" style="stop-color:#8b5cf6"/>
    </linearGradient>
    <linearGradient id="line" x1="0%" y1="0%" x2="100%" y2="0%">
      <stop offset="0%"   style="stop-color:#6366f1;stop-opacity:0"/>
      <stop offset="30%"  style="stop-color:#6366f1;stop-opacity:1"/>
      <stop offset="70%"  style="stop-color:#8b5cf6;stop-opacity:1"/>
      <stop offset="100%" style="stop-color:#8b5cf6;stop-opacity:0"/>
    </linearGradient>
  </defs>
  <rect width="960" height="160" rx="8" fill="url(#bg)"/>
  <g stroke="#ffffff" stroke-opacity="0.03" stroke-width="1">
    <line x1="0" y1="40"  x2="960" y2="40"/>
    <line x1="0" y1="80"  x2="960" y2="80"/>
    <line x1="0" y1="120" x2="960" y2="120"/>
    <line x1="240" y1="0" x2="240" y2="160"/>
    <line x1="480" y1="0" x2="480" y2="160"/>
    <line x1="720" y1="0" x2="720" y2="160"/>
  </g>
  <rect x="0" y="0" width="4" height="160" rx="2" fill="url(#accent)"/>
  <circle cx="40"  cy="80"  r="28" fill="#6366f1" fill-opacity="0.08"/>
  <circle cx="40"  cy="80"  r="18" fill="#6366f1" fill-opacity="0.10"/>
  <circle cx="40"  cy="80"  r="8"  fill="#6366f1" fill-opacity="0.20"/>
  <circle cx="880" cy="30"  r="50" fill="#8b5cf6" fill-opacity="0.05"/>
  <circle cx="920" cy="130" r="40" fill="#6366f1" fill-opacity="0.05"/>
  <text x="80" y="72"
        font-family="'Segoe UI','Helvetica Neue',Arial,sans-serif"
        font-size="36" font-weight="700" letter-spacing="-0.5"
        fill="#f8fafc">ShiftFestival</text>
  <text x="80" y="104"
        font-family="'Segoe UI','Helvetica Neue',Arial,sans-serif"
        font-size="16" font-weight="400" letter-spacing="0.5"
        fill="#94a3b8">Integration Platform — Desideriushogeschool · Groep 1</text>
  <line x1="80" y1="118" x2="880" y2="118" stroke="url(#line)" stroke-width="1"/>
  <rect x="80"  y="127" width="68" height="20" rx="4" fill="#0089d6" fill-opacity="0.20"/>
  <text x="114" y="141" font-family="'Segoe UI','Helvetica Neue',Arial,sans-serif"
        font-size="10" font-weight="600" text-anchor="middle" fill="#38bdf8">Azure VM</text>
  <rect x="156" y="127" width="60" height="20" rx="4" fill="#2496ed" fill-opacity="0.20"/>
  <text x="186" y="141" font-family="'Segoe UI','Helvetica Neue',Arial,sans-serif"
        font-size="10" font-weight="600" text-anchor="middle" fill="#60a5fa">Docker</text>
  <rect x="224" y="127" width="76" height="20" rx="4" fill="#ff6600" fill-opacity="0.20"/>
  <text x="262" y="141" font-family="'Segoe UI','Helvetica Neue',Arial,sans-serif"
        font-size="10" font-weight="600" text-anchor="middle" fill="#fb923c">RabbitMQ</text>
  <rect x="308" y="127" width="72" height="20" rx="4" fill="#00bfb3" fill-opacity="0.20"/>
  <text x="344" y="141" font-family="'Segoe UI','Helvetica Neue',Arial,sans-serif"
        font-size="10" font-weight="600" text-anchor="middle" fill="#2dd4bf">ELK Stack</text>
  <rect x="388" y="127" width="92" height="20" rx="4" fill="#6366f1" fill-opacity="0.20"/>
  <text x="434" y="141" font-family="'Segoe UI','Helvetica Neue',Arial,sans-serif"
        font-size="10" font-weight="600" text-anchor="middle" fill="#a5b4fc">2025 – 2026</text>
</svg>`
}

func main() {
	out := "../profile/badges"

	// Generate welcome banner.
	writeFile(filepath.Join(out, "banner.svg"), banner())

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

	fmt.Println("All assets generated.")
}
