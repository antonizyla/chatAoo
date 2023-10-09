<script lang="ts">
	import '../../app.css';

	import { onMount } from 'svelte';
	import { previousPage } from '$lib/stores/previousPageStore';

	onMount(async () => {
		const storedUser = localStorage.getItem('userID');
		if (!storedUser) {
			window.location.href = '/';
			localStorage.setItem('previousPage', window.location.href);
		}
	});
</script>

<slot />
