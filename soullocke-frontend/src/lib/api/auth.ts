import { createMutation, createQuery, useQueryClient } from '@tanstack/svelte-query';
import { checkAccessOptions, loginMutation, checkAccessQueryKey } from '$api/generated/@tanstack/svelte-query.gen';

export const getAuthMutation = () => {
	const queryClient = useQueryClient();

	return createMutation(() => {
		return {
			...loginMutation(),
			onSuccess: async (data, variables) => {
				void queryClient.invalidateQueries({
					queryKey: checkAccessQueryKey({ path: { lobbyId: variables.body.lobbyId } })
				});
				queryClient.setQueryData(checkAccessQueryKey({ path: { lobbyId: variables.body.lobbyId } }),
					{ canWrite: true });
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
			retry: false,
		};
	})
};
