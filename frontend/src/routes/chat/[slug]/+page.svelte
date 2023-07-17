<script lang="ts">
	import { onMount } from 'svelte';
	import type { PageData } from './$types';
	export let data: PageData;

	let userID: string | null = null;
	let userName: string | null = null;
	onMount(() => {
		// check if user exists on browser
		// if not, redirect to login page

		userID = localStorage.getItem('userID');
		userName = localStorage.getItem('userName');
		console.log(userID);
		if (!userID) {
			window.location.href = '/';
			// create store to store chat of user
			localStorage.setItem('chatId', data.chat.id);
		}
	});
</script>

<!-- svelte-ignore non-top-level-reactive-declaration -->
{#if data.exists}
	Chat exists with id {data.chat.id} and name '{data.chat.name}' with description '{data.chat
		.description}'
	<div>Chat goes here</div>
	<div class="">
		You are currently logged in as '{userName}' with id '{userID}'
	</div>
{:else}
	<h1>404</h1>
	<p>Chat does not exist</p>
{/if}
