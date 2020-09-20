package cli

import "github.com/kyleu/npn/npnweb"

func setContent() {
	setIcon()
}

func setIcon() {
	npnweb.IconContent = `<svg width="32px" height="32px" viewBox="-0 0 68 68" xmlns="http://www.w3.org/2000/svg">
				<g fill="none">
					<path id="logo-symbol" d="M 50.655 0 L 50.611 12.31 L 30.603 26.31 M 30.603 42.31 L 50.611 56.31 L 50.611 68.048 M 2 34.371 L 28.902 34.31 M 28.902 17.31 L 28.902 51.31 M 30.303 51.31 L 30.303 17.31 M 40.607 52.31 L 43.609 48.31 L 47.61 54.31 L 40.607 52.31 Z M 8.594 33.31 C 9.364 11.769 33.173 -0.86 51.451 10.577 C 69.728 22.014 68.766 48.94 49.718 59.043 C 31.449 68.734 9.332 55.971 8.594 35.31" style="stroke-width: 5px; paint-order: fill; stroke: rgb(135, 135, 135);"/>
				</g>
			</svg>`
}
