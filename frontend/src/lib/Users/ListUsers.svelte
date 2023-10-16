<script lang="ts">
	import Button from '$lib/svelte-components/components/Button/Button.svelte';
	import { onMount } from 'svelte';

	export let chat: string;

	let users: { user_id: string; user_name: string }[] = [];

	onMount(async () => {
		const res = await fetch(`http://localhost:8081/chats/${chat}/users`);
		users = await res.json();
	});

	import Modal from '$lib/svelte-components/components/Modal/Modal.svelte';
	let open: boolean = false;
</script>

<Button
	size="small"
	on:click={() => {
		open = !open;
	}}>{!open ? 'Show Users List' : 'Close Users List'}</Button
>

<Modal bind:showing={open}>
	<div class="text-md">Users in this chat:</div>
	<div class="text-sm p-2">
			{#each users as user}
					<div class="px-1.5 text-text/80 text-sm">{user.user_name}</div>
			{/each}
	</div>
</Modal>
