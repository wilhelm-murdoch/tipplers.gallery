<script lang="ts">
	import { Footer, Header } from '$components';
	import SvelteSeo from 'svelte-seo';

	const twitterMeta = {
		site: '@wilhelm',
		title: "Tippler's Gallery ~ Compendium",
		description: 'Quickly search & browse through thousands of cocktail recipes gathered from around the web.',
		image: 'https://tipplers.gallery/images/og-cover.png',
		imageAlt: "Tippler's Gallery"
	};

	const openGraphMeta = {
		title: "Tippler's Gallery ~ Compendium",
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

<div class="max-w-7xl mx-auto text-center pt-5 pb-10 md:pt-10">
	<h2 class="text-4xl md:text-6xl font-thin italic text-tipplers-secondary">Compendium</h2>
</div>

<slot />

<Footer />

<SvelteSeo title="Tippler's Gallery ~ Compendium" description="Quickly search & browse through thousands of cocktail recipes gathered from around the web." canonical="https://tipplers.gallery/compendium" twitter={twitterMeta} openGraph={openGraphMeta} />
