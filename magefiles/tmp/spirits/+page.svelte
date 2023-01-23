<script lang="ts">
	import { NavigationSub, NavigationFilter } from '$components';
	import type { NavigationItem } from '$components/Navigation/utils';
	import { onMount } from 'svelte';
	import { navigationItems } from '../utils';
	import { navigationItemsSpirits } from './utils';

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

<NavigationSub items={navigationItems} active="spirits" />
<NavigationFilter items={navigationItemsSpirits} active="Gin" wrapperClasses="-mb-1" />
<NavigationFilter items={alphaNumericFilterItems} active="a" />
