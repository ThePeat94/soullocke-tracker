<script lang="ts">
	import type { PageProps } from './$types';
	import { getLobbyQuery } from '$api/lobby';
	import PrimaryButton from '$components/buttons/PrimaryButton.svelte';
	import { getAccessQuery } from '$api/auth';
	import UnlockLobbyDialog from '$components/dialogs/UnlockLobbyDialog.svelte';

	let { params }: PageProps = $props();
	let unlockDialog = $state<ReturnType<typeof UnlockLobbyDialog>>();

	const lobbyQuery = getLobbyQuery(() => params.id);
	const accessCheck = getAccessQuery(() => params.id);

	const handleUnlockLobbyClick = (): void => {
		unlockDialog?.openDialog();
	};
</script>

{#if lobbyQuery.isLoading}
	Loading...
{/if}

{#if lobbyQuery.isError}
	{#if lobbyQuery.error.status === 404}
		<p>Lobby not found.</p>
	{:else}
		<p>An error occurred while fetching the lobby.</p>
	{/if}
{/if}

{#if lobbyQuery.isSuccess}
	<div class="mx-auto flex max-w-1/2 flex-col gap-4 mt-10">
		<div class="dark:bg-gray-800 rounded-2xl p-2">
			<h1 class="text-4xl text-center">{lobbyQuery.data.name}</h1>
		</div>
		{#if !(accessCheck.isSuccess)}
			<PrimaryButton onClick={handleUnlockLobbyClick} variant="filled">Join Lobby</PrimaryButton>
		{/if}
		<p>Game Edition: {lobbyQuery.data.gameEditionId}</p>
		<p>WIP!</p>

		<ul>
			TODO
			<li>- Add encounters</li>
			<li>- Add players</li>
			<li>- Add utility tools (optional)</li>
			<li>- Some nice UI/UX</li>
		</ul>
	</div>
{/if}

<svelte:head>
	<title>Lobby | {lobbyQuery.data?.name}</title>
</svelte:head>

<UnlockLobbyDialog bind:this={unlockDialog} lobbyId={params.id} />
