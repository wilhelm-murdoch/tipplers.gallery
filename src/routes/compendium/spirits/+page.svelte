<script lang="ts">
	import CardSlim from '$components/Cards/CardSlim.svelte';
	import Footer from '$components/Footer/Footer.svelte';
	import Header from '$components/Header/Header.svelte';
	import NavigationSub from '$components/Navigation/NavigationSub.svelte';
	import { navigationItems } from '../utils';
	import { onMount } from 'svelte';
	import SvelteSeo from 'svelte-seo';
	import { fade } from 'svelte/transition';

	let data: any;

	const load = async () => {
		const res = await fetch(`/api/cocktails-thebar.json`);
		let cocktails = await res.json();

		return { cocktails };
	};

	const chars: string = '#abcdefghijklmnopqrstuvwxyz';
	let charsNav: any[] = [];

	onMount(async () => {
		data = await load();
		chars.split('').forEach((c) => {
			charsNav.push({
				char: c.toUpperCase(),
				cocktails: data.cocktails.filter((v: any) => v.name[0].toLowerCase() === c)
			});
		});

		charsNav = charsNav;
	});

	const twitterMeta = {
		site: '@wilhelm',
		title: "Tippler's Gallery ~ Compendium / Spirits",
		description: 'Quickly search & browse through thousands of cocktail recipes gathered from around the web.',
		image: 'https://tipplers.gallery/images/og-cover.png',
		imageAlt: "Tippler's Gallery"
	};

	const openGraphMeta = {
		title: "Tippler's Gallery ~ Compendium / Spirits",
		description: 'Quickly search & browse through thousands of cocktail recipes gathered from around the web.',
		type: 'article',
		url: 'https://tipplers.gallery/compendium',
		images: [
			{
				url: 'https://tipplers.gallery/images/og-cover.png',
				alt: "Tippler's Gallery"
			}
		]
	};
</script>

<div class="px-2">
	<Header />
</div>

<div class="max-w-7xl mx-auto text-center py-10">
	<h2 class="text-6xl font-thin italic text-tipplers-secondary">The Compendium</h2>
</div>

<NavigationSub items={navigationItems} active="spirits" />

<div class="bg-white border-y">
	<div class="max-w-7xl mx-auto">
		<nav class="flex justify-between py-5 bg-white overflow-x-auto">
			<a class="rounded-md px-4 py-2.5 hover:bg-tipplers-primary hover:text-white" href="/" title="">Gin</a>
			<a class="rounded-md px-4 py-2.5 bg-tipplers-primary text-white" href="/" title="">Whiskey</a>
			<a class="rounded-md px-4 py-2.5 hover:bg-tipplers-primary hover:text-white" href="/" title="">Rum</a>
			<a class="rounded-md px-4 py-2.5 hover:bg-tipplers-primary hover:text-white" href="/" title="">Cognac</a>
			<a class="rounded-md px-4 py-2.5 hover:bg-tipplers-primary hover:text-white" href="/" title="">Absinthe</a>
			<a class="rounded-md px-4 py-2.5 hover:bg-tipplers-primary hover:text-white" href="/" title="">Tequila</a>
			<a class="rounded-md px-4 py-2.5 hover:bg-tipplers-primary hover:text-white" href="/" title="">Mezcal</a>
			<a class="rounded-md px-4 py-2.5 hover:bg-tipplers-primary hover:text-white" href="/" title="">Vodka</a>
		</nav>
	</div>
</div>

<div class="bg-white border-b">
	<div class="max-w-7xl mx-auto">
		<nav class="flex justify-between py-5 bg-white overflow-x-auto">
			{#each charsNav as item, i}
				{#if item.cocktails.length}
					<a class="rounded-md px-4 py-2.5 hover:bg-tipplers-primary hover:text-white" href="#{item.char}" title={item.char}>
						{item.char}
					</a>
				{:else}
					<span class="px-4 py-2.5 text-tipplers-secondary">{item.char}</span>
				{/if}
			{/each}
		</nav>
	</div>
</div>

<div class="max-w-7xl mx-auto px-2">
	{#each charsNav as item, i}
		{#if item.cocktails.length}
			<section>
				<h3 id={item.char} class="py-8 text-2xl font-semibold font-serif border-b border-slate-50 bg-tipplers-lightest/25">
					<span class="underline">{item.char}</span>
					<span class="bg-white px-2 py-1 rounded shadow-sm text-xs font-sans align-top">
						{item.cocktails.length}
					</span>
				</h3>
				<div class="grid grid-cols-2 gap-6 md:grid-cols-4 lg:grid-cols-6">
					{#each item.cocktails as cocktail, i}
						<div in:fade={{ delay: i * 100 }}>
							<CardSlim name={cocktail.name} src={cocktail.images[0].relative_path} from={cocktail.meta.sourced_from} attribution={cocktail.images[0].attribution} url={cocktail.meta.source_url} />
						</div>
					{/each}
				</div>
			</section>
		{/if}
	{/each}
</div>

<Footer />

<SvelteSeo title="Tippler's Gallery ~ Compendium / Spirits" description="Quickly search & browse through thousands of cocktail recipes gathered from around the web." canonical="https://tipplers.gallery/compendium/spirits" twitter={twitterMeta} openGraph={openGraphMeta} />
