module.exports = {
	content: [
		'./src/**/*.{html,js,svelte,ts}'
	],
	theme: {
		extend: {
			colors: {
				'tipplers-primary': 'rgb(87, 155, 177)',
				'tipplers-secondary': 'rgb(201 188 165)',
				'tipplers-lighter': 'rgb(236, 232, 221)',
				'tipplers-lightest': 'rgb(248, 244, 234)',
			}
		}
	},
	plugins: [
		require('@tailwindcss/forms'),
		require('@tailwindcss/typography')
	]
};