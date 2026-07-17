import { createMutation, createQuery, useQueryClient } from '@tanstack/svelte-query';
import { checkAccessOptions, loginMutation, checkAccessQueryKey } from '$api/generated/@tanstack/svelte-query.gen';

export const getAuthMutation = () => {
	const queryClient = useQueryClient();

	return createMutation(() => {
		return {
			...loginMutation(),
			onSuccess: async (data, variables) => {
				await queryClient.invalidateQueries({
					queryKey: checkAccessQueryKey({path: {lobbyId: variables.body.lobbyId}})
				})
			}
		};
	})
};

export const getAccessQuery = (getId: () => string) => {
	return createQuery(() => {
		const id = getId();
		return {
			...checkAccessOptions({ path: { lobbyId: id }}),
			enabled: id.length > 0,
			refetchInterval: 5000,
		};
	})
};
