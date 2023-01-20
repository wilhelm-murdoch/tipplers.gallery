<script lang="ts">
	import SvelteSeo from 'svelte-seo';

	const twitterMeta = {
		site: '@wilhelm',
		title: "Tippler's Gallery &middot; A compendium of the finest cocktail recipes.",
		description: 'Quickly search & browse through thousands of cocktail recipes gathered from around the web.',
		image: 'https://tipplers.gallery/images/og-cover.png',
		imageAlt: "Tippler's Gallery"
	};

	const openGraphMeta = {
		title: "Tippler's Gallery &middot; A compendium of the finest cocktail recipes.",
		description: 'Quickly search & browse through thousands of cocktail recipes gathered from around the web.',
		type: 'article',
		url: 'https://tipplers.gallery',
		images: [
			{
				url: 'https://tipplers.gallery/images/og-cover.png',
				alt: "Tippler's Gallery"
			}
		]
	};
</script>

<div class="flex flex-col h-screen bg-tipplers-lightest/50 overflow-auto p-4 select-none">
	<div class="m-auto">
		<h1 class="mb-4 text-3xl md:text-6xl h-16 md:h-32 text-center">
			<object class="shake-animation h-full inline-block w-16 md:w-32" data="/images/cheers.svg" type="image/svg+xml" title="" />
			<span class="font-bold text-center underline font-serif align-middle">Tippler's Gallery</span>
		</h1>
	</div>

	<footer class="text-center text-sm text-slate-400 mt-16 mb-4">
		Made with
		<svg class="w-6 h-6 text-red-600 inline-block heartbeat-animation" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
			<path fill-rule="evenodd" d="M3.172 5.172a4 4 0 015.656 0L10 6.343l1.172-1.171a4 4 0 115.656 5.656L10 17.657l-6.828-6.829a4 4 0 010-5.656z" clip-rule="evenodd" />
		</svg>
		by
		<a href="https://wilhelm.codes" title="It me." class="font-semibold border-b hover:text-slate-400">wilhelm</a>
		<img class="inline-block h-6 w-6 rounded-full border-2 border-slate-300 hover:border-slate-400 jello-animation" src="https://s.gravatar.com/avatar/7cc6bee5975704b299bd7355e8ba9356?s=40" alt="It me." />
		<!-- svelte-ignore missing-declaration -->
		<div class="text-center text-sm text-tipplers-secondary mt-2.5 font-mono">{__BUILD_NUMBER__}</div>
	</footer>
</div>

<SvelteSeo title="Tippler's Gallery &middot; A compendium of the finest cocktail recipes." description="Quickly search & browse through thousands of cocktail recipes gathered from around the web." canonical="https://tipplers.gallery" twitter={twitterMeta} openGraph={openGraphMeta} />

<style>
	.shake-animation,
	.shake-animation:hover {
		-webkit-animation: shake-animation 0.8s ease-in-out 250ms 2 both;
		animation: shake-animation 0.8s ease-in-out 250ms 2 both;
	}

	@-webkit-keyframes shake-animation {
		0%,
		100% {
			-webkit-transform: rotate(0deg);
			transform: rotate(0deg);
			-webkit-transform-origin: 50% 100%;
			transform-origin: 50% 100%;
		}
		10% {
			-webkit-transform: rotate(2deg);
			transform: rotate(2deg);
		}
		20%,
		40%,
		60% {
			-webkit-transform: rotate(-4deg);
			transform: rotate(-4deg);
		}
		30%,
		50%,
		70% {
			-webkit-transform: rotate(4deg);
			transform: rotate(4deg);
		}
		80% {
			-webkit-transform: rotate(-2deg);
			transform: rotate(-2deg);
		}
		90% {
			-webkit-transform: rotate(2deg);
			transform: rotate(2deg);
		}
	}
	@keyframes shake-animation {
		0%,
		100% {
			-webkit-transform: rotate(0deg);
			transform: rotate(0deg);
			-webkit-transform-origin: 50% 100%;
			transform-origin: 50% 100%;
		}
		10% {
			-webkit-transform: rotate(2deg);
			transform: rotate(2deg);
		}
		20%,
		40%,
		60% {
			-webkit-transform: rotate(-4deg);
			transform: rotate(-4deg);
		}
		30%,
		50%,
		70% {
			-webkit-transform: rotate(4deg);
			transform: rotate(4deg);
		}
		80% {
			-webkit-transform: rotate(-2deg);
			transform: rotate(-2deg);
		}
		90% {
			-webkit-transform: rotate(2deg);
			transform: rotate(2deg);
		}
	}

	.heartbeat-animation:hover {
		-webkit-animation: heartbeat-animation 1.5s ease-in-out 2 both;
		animation: heartbeat-animation 1.5s ease-in-out 2 both;
	}

	@-webkit-keyframes heartbeat-animation {
		from {
			-webkit-transform: scale(1);
			transform: scale(1);
			-webkit-transform-origin: center center;
			transform-origin: center center;
			-webkit-animation-timing-function: ease-out;
			animation-timing-function: ease-out;
		}

		10% {
			-webkit-transform: scale(0.91);
			transform: scale(0.91);
			-webkit-animation-timing-function: ease-in;
			animation-timing-function: ease-in;
		}

		17% {
			-webkit-transform: scale(0.98);
			transform: scale(0.98);
			-webkit-animation-timing-function: ease-out;
			animation-timing-function: ease-out;
		}

		33% {
			-webkit-transform: scale(0.87);
			transform: scale(0.87);
			-webkit-animation-timing-function: ease-in;
			animation-timing-function: ease-in;
		}

		45% {
			-webkit-transform: scale(1);
			transform: scale(1);
			-webkit-animation-timing-function: ease-out;
			animation-timing-function: ease-out;
		}
	}

	@keyframes heartbeat-animation {
		from {
			-webkit-transform: scale(1);
			transform: scale(1);
			-webkit-transform-origin: center center;
			transform-origin: center center;
			-webkit-animation-timing-function: ease-out;
			animation-timing-function: ease-out;
		}

		10% {
			-webkit-transform: scale(0.91);
			transform: scale(0.91);
			-webkit-animation-timing-function: ease-in;
			animation-timing-function: ease-in;
		}

		17% {
			-webkit-transform: scale(0.98);
			transform: scale(0.98);
			-webkit-animation-timing-function: ease-out;
			animation-timing-function: ease-out;
		}

		33% {
			-webkit-transform: scale(0.87);
			transform: scale(0.87);
			-webkit-animation-timing-function: ease-in;
			animation-timing-function: ease-in;
		}

		45% {
			-webkit-transform: scale(1);
			transform: scale(1);
			-webkit-animation-timing-function: ease-out;
			animation-timing-function: ease-out;
		}
	}

	.jello-animation:hover {
		-webkit-animation: jello-animation 0.9s both;
		animation: jello-animation 0.9s both;
	}

	@-webkit-keyframes jello-animation {
		0% {
			-webkit-transform: scale3d(1, 1, 1);
			transform: scale3d(1, 1, 1);
		}

		30% {
			-webkit-transform: scale3d(1.25, 0.75, 1);
			transform: scale3d(1.25, 0.75, 1);
		}

		40% {
			-webkit-transform: scale3d(0.75, 1.25, 1);
			transform: scale3d(0.75, 1.25, 1);
		}

		50% {
			-webkit-transform: scale3d(1.15, 0.85, 1);
			transform: scale3d(1.15, 0.85, 1);
		}

		65% {
			-webkit-transform: scale3d(0.95, 1.05, 1);
			transform: scale3d(0.95, 1.05, 1);
		}

		75% {
			-webkit-transform: scale3d(1.05, 0.95, 1);
			transform: scale3d(1.05, 0.95, 1);
		}

		100% {
			-webkit-transform: scale3d(1, 1, 1);
			transform: scale3d(1, 1, 1);
		}
	}

	@keyframes jello-animation {
		0% {
			-webkit-transform: scale3d(1, 1, 1);
			transform: scale3d(1, 1, 1);
		}

		30% {
			-webkit-transform: scale3d(1.25, 0.75, 1);
			transform: scale3d(1.25, 0.75, 1);
		}

		40% {
			-webkit-transform: scale3d(0.75, 1.25, 1);
			transform: scale3d(0.75, 1.25, 1);
		}

		50% {
			-webkit-transform: scale3d(1.15, 0.85, 1);
			transform: scale3d(1.15, 0.85, 1);
		}

		65% {
			-webkit-transform: scale3d(0.95, 1.05, 1);
			transform: scale3d(0.95, 1.05, 1);
		}

		75% {
			-webkit-transform: scale3d(1.05, 0.95, 1);
			transform: scale3d(1.05, 0.95, 1);
		}

		100% {
			-webkit-transform: scale3d(1, 1, 1);
			transform: scale3d(1, 1, 1);
		}
	}
</style>
