<script lang="ts">
	import Button from '$lib/components/Button/Button.svelte';
	import { onMount } from 'svelte';

	export let chat: string;

	let open: boolean = false;
	function toggle() {
		open = !open;
	}

	let users: { user_id: string; user_name: string }[] = [];

	onMount(async () => {
		const res = await fetch(`http://localhost:8081/chats/${chat}/users`);
		users = await res.json();
		console.log(users);
	});
</script>

<div class="flex flex-col gap-2 w-fit p-2">
	<Button size="medium" on:click={toggle}>{!open ? 'Show Users List' : 'Close Users List'}</Button>
	{#if open}
		<div class="text-md">Users in this chat:</div>
		<div class="text-sm">
			{#each users as user}
				<div class="">{user.user_name}</div>
			{/each}
		</div>
	{/if}
</div>
