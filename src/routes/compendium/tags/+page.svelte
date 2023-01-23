<script lang="ts">
	import Footer from '$components/Footer/Footer.svelte';
	import Header from '$components/Header/Header.svelte';
	import NavigationSub from '$components/Navigation/NavigationSub.svelte';
	import { navigationItems } from '../utils';

	export let data: any;
</script>

<div class="px-2">
	<Header />
</div>

<div class="max-w-7xl mx-auto text-center py-10">
	<h2 class="text-6xl font-thin italic">The Compendium</h2>
</div>

<NavigationSub items={navigationItems} active="tags" />

<Footer />
