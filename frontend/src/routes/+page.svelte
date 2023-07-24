<script lang="ts">
	import { onMount } from 'svelte';
	import type { ActionData } from './$types';
	import { enhance } from '$app/forms';
	export let form: ActionData;

	let userName: string | null = '';
	let userID: string | null = '';
	let mounted = false;
	onMount(() => {
		// check local storage for a user tag
		mounted = true;
		userName = localStorage.getItem('userName');
		userID = localStorage.getItem('userID');
	});

	$: {
		if (form?.user && mounted) {
			localStorage.setItem('userName', form.user.name);
			localStorage.setItem('userID', form.user.user_id);
			userName = form.user.username;
			const chatId = localStorage.getItem('chatId');
			if (chatId) {
				window.location.href = `/chat/${chatId}`;
				localStorage.removeItem('chatId');
			}
		}
	}
</script>

<div class="p-12 bg-red-100 w-fit mx-auto mt-20">
	{#if userName || form?.user}
		<div class="">
			You are currently signed in as <p class="bg-blue-100 inline">{userName || form?.user.name}</p>
		</div>
		<p>
			Click <a class="bg-blue-100 border-blue-400 border-2 p-0.5" href="/chat">here</a> to go to create
			or join a chat
		</p>
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
