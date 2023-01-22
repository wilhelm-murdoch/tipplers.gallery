<script lang="ts">
	import Footer from '$components/Footer/Footer.svelte';
	import LogoLarge from '$components/Logo/LogoLarge.svelte';
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
		<LogoLarge />
	</div>

	<Footer />
</div>

<SvelteSeo title="Tippler's Gallery &middot; A compendium of the finest cocktail recipes." description="Quickly search & browse through thousands of cocktail recipes gathered from around the web." canonical="https://tipplers.gallery" twitter={twitterMeta} openGraph={openGraphMeta} />
