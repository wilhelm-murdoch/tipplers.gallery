<script lang="ts">
	import { NavigationSub, NavigationFilter } from '$components';
	import type { NavigationItem } from '$components/Navigation/utils';
	import { onMount } from 'svelte';
	import { navigationItems } from '../utils';

	let alphaNumericFilterItems: NavigationItem[] = [];

	onMount(() => {
		alphaNumericFilterItems = '#abcdefghijklmnopqrstuvwxyz'.split('').map((c) => {
			return {
				name: c.toUpperCase(),
				slug: c,
				path: '',
				title: 'Filter cocktails starting with "' + c + '".',
				count: 0
			};
		});
	});
</script>

<NavigationSub items={navigationItems} active="tags" />
<NavigationFilter items={alphaNumericFilterItems} active="a" />
