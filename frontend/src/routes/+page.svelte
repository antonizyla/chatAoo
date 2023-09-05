<script lang="ts">
	import { onMount } from 'svelte';
	import type { ActionData } from './$types';
	import { enhance } from '$app/forms';
	import Button from '$lib/components/Button/Button.svelte';
	export let form: ActionData;

	let userName: string | null = '';
	let userID: string | null = '';
	let mounted = false;
	onMount(async () => {
		// check local storage for a user tag
		mounted = true;
		userName = localStorage.getItem('userName');
		userID = localStorage.getItem('userID');
	});

	async function linkUser(chatID: string, userID: string) {
		console.log('linking user');
		const link = await fetch(`http://localhost:8081/chats/link`, {
			method: 'POST',
			body: JSON.stringify({ chat_id: chatID, user_id: userID })
		}).then((res) => {
			res.status, res.statusText;
		});
		console.log(link);
		// redirect to chatpage
		window.location.href = `/chat/${chatID}`;
		localStorage.removeItem('chatId');
	}

	$: {
		if (form?.user && mounted) {
			localStorage.setItem('userName', form.user.name);
			localStorage.setItem('userID', form.user.user_id);
			userName = form.user.username;
			const chatID = localStorage.getItem('chatId');
			if (chatID) {
				linkUser(chatID, form.user.user_id);
			}
		}
	}
</script>

<div class="p-12 bg-red-100 w-fit mx-auto mt-20">
	{#if userName || form?.user}
		<div class="p-2">
			You are currently signed in as <p class="bg-blue-100 inline">{userName || form?.user.name}</p>
		</div>
		<div class="flex flex-row p-2 gap-2">
			<a href="/account"><Button size="small">Edit Account Details</Button></a>
			<a href="/chat"><Button size="small">Create or Enter a chat</Button></a>
		</div>
	{:else}
		<p>No User Identifier has been found in your browser</p>
		<p>Create one below</p>
		<form action="?/createUser" method="POST" use:enhance>
			<label for="name">User Identifier</label>
			<input type="text" name="name" id="user" placeholder="Enter a user identifier" required />
			<button type="submit" on:click={() => {}}>Create User</button>
		</form>
	{/if}
</div>
