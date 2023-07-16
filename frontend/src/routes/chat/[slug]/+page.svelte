<script lang="ts">
	import { onMount } from 'svelte';

	// page load
	import type { ActionData, PageData } from './$types';
	export let data: PageData;

	export let form: ActionData;

	let mounted = false;
	$: {
		if (form?.exists && mounted) {
			localStorage.setItem(`chat-${data.chat.id}-user`, form.user);
		}
	}

	let userExists = false;
	let user: string | null;
	onMount(() => {
		mounted = true;
		// check local storage for a user with this chat id
		user = localStorage.getItem(`chat-${data.chat.id}-user`);
		if (user) {
			userExists = true;
		}
	});
</script>

<!-- svelte-ignore non-top-level-reactive-declaration -->
{#if data.exists}
	<h1>Chat with Name '{data.chat.name}' has been Found</h1>
	{#if data.chat.description}
		<p>Chat Description: {data.chat.description}</p>
	{/if}
	{#if !userExists}
		<p>Enter a display name to enter the chat</p>
		<form action="?/generateUser" method="POST">
			<input type="text" name="username" placeholder="Display Name" />
			<input type="text" hidden name="chatID" bind:value={data.chat.id} />
			<button type="submit">Join Chat</button>
		</form>
	{:else}
		Your Current User: {user}
		<div class="">show chats</div>
	{/if}
{:else}
	<h1>404</h1>
	<p>Chat does not exist</p>
{/if}
