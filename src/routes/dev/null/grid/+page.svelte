<script lang="ts">
	import { lazy } from '$lib/utils/lazy';
	import { onMount } from 'svelte';
	import { fade } from 'svelte/transition';

	export let data: any;

	const sleep = (ms: number) => new Promise((r) => setTimeout(r, ms));

	const getRandomElements = (arr: any[], n: number) => {
		var result = new Array(n);
		var len = arr.length;
		var taken = new Array(len);

		while (n--) {
			var x = Math.floor(Math.random() * len);
			result[n] = arr[x in taken ? taken[x] : x];
			taken[x] = --len in taken ? taken[len] : len;
		}
		return result;
	};

	const setRandomInterval = (intervalFunction: any, minDelay: number, maxDelay: number) => {
		let timeout: any;

		const runInterval = () => {
			const timeoutFunction = () => {
				intervalFunction();
				runInterval();
			};

			const delay = Math.floor(Math.random() * (maxDelay - minDelay + 1)) + minDelay;

			timeout = setTimeout(timeoutFunction, delay);
		};

		runInterval();

		return {
			clear() {
				clearTimeout(timeout);
			}
		};
	};

	const getRandomHEX = () => {
		return Math.floor(Math.random()*16777215).toString(16);
	}

	let limitRandom = 10;
	let limitCocktails = 100;
	let minDelay = 1000;
	let maxDelay = 4000;
	let cocktails: any[] = [];

	onMount(() => {
		cocktails = getRandomElements(data.cocktails, limitCocktails);

		const interval = setRandomInterval(() => {
				for (var i = 0; i < Math.floor(Math.random() * limitRandom); i++) {
					cocktails[Math.floor(Math.random() * limitCocktails)] = data.cocktails[Math.floor(Math.random() * data.cocktails.length)]
				}
			},
			minDelay,
			maxDelay
		);
	});
</script>

<div class="grid grid-cols-4 md:grid-cols-5 lg:grid-cols-10 gap-0 select-none">
	{#each cocktails as cocktail, i}
		<div class="relative overflow-hidden aspect-square cursor-pointer" transition:fade={{ delay: i*50}}>
			<div class="hover:scale-110 ease-in-out duration-500">
				{#key cocktail.cid}
					<img src={`/images/cocktails/${cocktail.cid}/${cocktail.i}`} class="absolute inset-0" alt="" transition:fade={{ duration: 1000 }}/>
				{/key}
			</div>
		</div>
	{/each}
</div>