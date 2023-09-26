<script lang="ts">
	import '../../app.css';

	import { onMount } from 'svelte';
	import { previousPage } from '$lib/stores/previousPageStore';

	onMount(async () => {
		const storedUser = localStorage.getItem('userID');
		if (!storedUser) {
			location.href = '/';
			$previousPage = location.href;
		}
	});
</script>

<slot />
