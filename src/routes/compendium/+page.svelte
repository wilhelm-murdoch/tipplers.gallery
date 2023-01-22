<script lang="ts">
	import CardSlim from '$components/Cards/CardSlim.svelte';
	import Footer from '$components/Footer/Footer.svelte';
	import Header from '$components/Header/Header.svelte';
	import { onMount } from 'svelte';
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
</script>

<Header />

<div class="max-w-7xl mx-auto text-center py-10">
	<h2 class="text-6xl font-thin italic">The Compendium</h2>
</div>
{#if charsNav.length}
	<div class="bg-white border-y">
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
						<span class="bg-white px-2 py-1 rounded shadow-sm text-xs font-normal font-sans align-top">
							<strong>{item.cocktails.length}</strong>
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
{/if}

<Footer />
