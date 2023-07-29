<script lang="ts">
	import { users, type message } from '$lib/messagesStores';
	import { timeDisplay } from '$lib/timeDisplay';

	export let message: message;
	import MessageSettings from './MessageSettings.svelte';
</script>

<div class="flex flex-col bg-red-50 p-2 m-4">
	<div class="bg-blue-50">
		<div class="font-semibold text-lg">
			{$users[message.user_id]}
		</div>
		{#if !message.deleted}
			<div class="content">
				{message.body}
			</div>
		{:else}
			<div class="content">Message has been deleted</div>
		{/if}
	</div>
	<div class="flex flex-row justify-end bg-blue-100 gap-2">
		{timeDisplay(message.created_at)}
		{#if message.updated_at != message.created_at}
			{#if message.deleted}
				(deleted {timeDisplay(message.updated_at)})
			{:else}
				(edited {timeDisplay(message.updated_at)})
			{/if}
		{/if}
		<MessageSettings {message} />
	</div>
</div>
