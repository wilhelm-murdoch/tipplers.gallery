<script lang="ts">
	import CardCollection from '$components/Cards/CardCollection.svelte';
	import CardResult from '$components/Cards/CardResult.svelte';
	import CardSlim from '$components/Cards/CardSlim.svelte';
	import CardSlimLoading from '$components/Cards/CardSlimLoading.svelte';
	import CardSummary from '$components/Cards/CardSummary.svelte';
	import { onMount } from 'svelte';

	export let data: any;

	let randomSummaryCardCocktail: number;

	const getRandomCocktail = function () {
		return Math.floor(Math.random() * data.cocktails.length);
	};

	onMount(() => {
		randomSummaryCardCocktail = getRandomCocktail();
	});
</script>

<div class="flex flex-col h-screen p-4 max-w-7xl mx-auto">
	<h3 class="font-serif font-bold text-2xl border-b border-tipplers-secondary my-4">Design Elements</h3>
	<navigation class="prose">
		<ul>
			<li><a href="#summary-cards">Summary Cards</a></li>
			<li><a href="#slim-card-loading">Slim Card Loading</a></li>
			<li><a href="#slim-cards">Slim Cards</a></li>
			<li><a href="#result-cards">Result Cards</a></li>
			<li><a href="#collection-cards">Collection Cards</a></li>
			<li><a href="/dev/null/grid">Crossfade Grid</a></li>
		</ul>
	</navigation>

	<div class="container mx-auto max-w-7xl">
		<h3 class="font-serif font-bold text-2xl border-b border-tipplers-secondary my-4" id="summary-cards">Summary Cards</h3>
		{#if randomSummaryCardCocktail}
			<CardSummary cocktail={data.cocktails[randomSummaryCardCocktail]} />
		{/if}
	</div>

	<div class="container mx-auto max-w-7xl">
		<h3 class="font-serif font-bold text-2xl border-b border-tipplers-secondary my-4" id="slim-card-loading">Slim Card Loading</h3>
		<div class="grid grid-cols-2 gap-6 md:grid-cols-4 lg:grid-cols-6">
			<CardSlimLoading/>
			<CardSlimLoading/>
			<CardSlimLoading/>
			<CardSlimLoading/>
			<CardSlimLoading/>
			<CardSlimLoading/>
	</div>
		<h3 class="font-serif font-bold text-2xl border-b border-tipplers-secondary my-4" id="slim-cards">Slim Cards</h3>
		<div class="grid grid-cols-2 gap-6 md:grid-cols-4 lg:grid-cols-6">
			{#each data.cocktails.slice(0, 12) as cocktail, i}
				<CardSlim name={cocktail.name} src={cocktail.images[0].relative_path} from={cocktail.meta.sourced_from} attribution={cocktail.images[0].attribution} url={cocktail.meta.source_url} />
			{/each}
		</div>
	</div>

	<div class="container mx-auto max-w-7xl">
		<h3 class="font-serif font-bold text-2xl border-b border-tipplers-secondary my-4" id="result-cards">Result Cards</h3>
		<div class="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3">
			{#each data.cocktails.slice(0, 12) as cocktail}
				<CardResult {cocktail} />
			{/each}
		</div>
	</div>

	<div class="container mx-auto max-w-7xl">
		<h3 class="font-serif font-bold text-2xl border-b border-tipplers-secondary my-4" id="collection-cards">Collection Cards</h3>
		<div class="grid grid-cols-1 gap-6 md:grid-cols-2 pb-4">
			<CardCollection cocktails={data.cocktails.slice(0, 6)} />
			<CardCollection cocktails={data.cocktails.slice(0, 6)} />
			<CardCollection cocktails={data.cocktails.slice(0, 6)} />
			<CardCollection cocktails={data.cocktails.slice(0, 6)} />
		</div>
	</div>
</div>
