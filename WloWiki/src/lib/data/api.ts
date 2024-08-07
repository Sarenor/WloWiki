import type { State } from '@vincjo/datatables/remote';

export const reload = async (state: State, type: String, providedFetch?: any) => {
	const response = await providedFetch ? providedFetch(`/internalapi/${type}?${getParams(state)}`) : fetch(`/internalapi/${type}?${getParams(state)}`);
	return response.json();
};

const getParams = (state: State) => {
	const { pageNumber, rowsPerPage, sort, filters, search } = state;

	let params = `_page=${pageNumber}`;

	if (rowsPerPage) {
		params += `&_limit=${rowsPerPage}`;
	}
	if (sort) {
		params += `&_sort=${sort.orderBy}&_order=${sort.direction}`;
	}

	if (filters) {
		params += filters.map(({ filterBy, value }) => `&${filterBy}=${value}`).join();
	}
	// console.log(params);
	if (search) {
		params += `&q=${search}`;
	}
	return params;
};