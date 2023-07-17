<script lang="ts">
	import { onMount } from 'svelte';
	import type { ActionData } from './$types';
	import { enhance } from '$app/forms';
	export let form: ActionData;

	let userID: string | null;
	onMount(() => {
		userID = localStorage.getItem('userID');
	});
</script>

<h1>Welcome to Chat App</h1>
{#if !form?.success}
	<h2>Create a New Chat</h2>
	<form action="?/createChat" method="POST" use:enhance>
		<input type="text" name="name" placeholder="Chat Name" />
		<input type="text" name="description" placeholder="Chat Description" />
		<input type="text" name="userID" bind:value={userID} hidden />
		<button type="submit"> Create Chat </button>
	</form>
	<h2>Join an existing Chat</h2>
	<form action="?/joinChat" method="POST" use:enhance>
		<input type="text" name="uuid" placeholder="unique chat id" />
		<input type="text" hidden bind:value={userID} name="userID" />
		<button type="submit">Join Chat</button>
	</form>
{:else}
	Successfully joined Chat with name {form?.name} and url: /chat/{form.uuid}
	<button><a href={`/chat/${form.uuid}`}>Launch Chat</a></button>
{/if}
