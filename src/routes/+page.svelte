<div class="flex flex-col h-screen bg-tipplers-lightest/50 overflow-auto p-4">
	<div class="m-auto max-w-3xl">
		<h1 class="text-3xl font-bold text-center underline font-serif mb-4">Tippler's Gallery</h1>
		<h2 class="font-bold tracking-tight text-center text-6xl">A <span class="text-tipplers-primary">compendium</span> of the <i class="font-serif text-tipplers-primary/80 font-normal">finest</i> cocktail recipes.</h2>
		<p class="mt-6 text-lg leading-8 text-slate-800 text-center">This project is currently under heavy development and <span class="italic underline font-serif">coming soon</span>. In the mean time, feel free to download the current library of delicious boozy concoctions below.</p>
		<div class="mt-8 flex gap-x-4 justify-center">
			<a href="/cocktails.json" class="inline-block rounded-lg bg-gradient-to-tr from-tipplers-primary to-blue-500 px-4 py-1.5 text-base font-semibold leading-7 text-white shadow-md hover:bg-tipplers-primary/90">
				Download
				<svg class="inline-block h-6 w-6" fill="none" stroke="currentColor" stroke-width="1.5" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" aria-hidden="true">
					<path stroke-linecap="round" stroke-linejoin="round" d="M12 9.75v6.75m0 0l-3-3m3 3l3-3m-8.25 6a4.5 4.5 0 01-1.41-8.775 5.25 5.25 0 0110.233-2.33 3 3 0 013.758 3.848A3.752 3.752 0 0118 19.5H6.75z" />
				</svg>
			</a>
		</div>
	</div>

	<footer class="text-center text-sm text-slate-400 mt-16 mb-4">
		Made with
		<svg class="w-6 h-6 text-red-600 inline-block heartbeat-animation" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M3.172 5.172a4 4 0 015.656 0L10 6.343l1.172-1.171a4 4 0 115.656 5.656L10 17.657l-6.828-6.829a4 4 0 010-5.656z" clip-rule="evenodd" /></svg>
		by
		<a href="https://wilhelm.codes" title="It me." class="font-semibold border-b hover:text-slate-400">wilhelm</a>
		<img class="inline-block h-6 w-6 rounded-full border-2 border-slate-300 hover:border-slate-400 jello-animation" src="https://s.gravatar.com/avatar/7cc6bee5975704b299bd7355e8ba9356?s=40" alt="It me." />
		<!-- svelte-ignore missing-declaration -->
		<div class="text-center text-sm text-tipplers-secondary mt-2.5 font-mono">{__BUILD_NUMBER__}</div>
	</footer>
</div>

<style>
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
