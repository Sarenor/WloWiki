import { reload } from '$lib/data/api';
import type { State } from '@vincjo/datatables/local';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params, fetch }) => {
	return {
        materials: await reload({pageNumber: 0} as State, "materials", fetch)
	};
};