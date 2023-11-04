import { defineType } from "sanity";

export default {
    name: 'plot',
    type: 'document',
    title: 'Plot',
    fields: [
		defineType({
			title: 'Project Name',
			name: 'project_name',
			type: 'string',
		}),
		defineType({
			title: 'Images',
			name: 'images',
			type: 'array',
			of: [
				{
					type: 'image',
				},
			],
		}),
		defineType({
			type: 'object',
			name: 'lower_limit',
			title: 'Lower Limit',
			fields: [
				{type: 'number', name: 'lower_limit_price', title: 'Lower Limit Price'},
				{
					type: 'string',
					name: 'denomination',
					title: 'Denomination',
					options: {
						list: ['Cr', 'L'],
					},
				},
			],
		}),
		defineType({
			type: 'object',
			name: 'upper_limit',
			title: 'Upper Limit',
			fields: [
				{type: 'number', name: 'upper_limit_price', title: 'Upper Limit Price'},
				{
					type: 'string',
					name: 'denomination',
					title: 'Denomination',
					options: {
						list: ['Cr', 'L'],
					},
				},
			],
		}),
		defineType({
			type: 'number',
			name: 'lower_limit_size',
			title: 'Lower Limit Size (sq. ft)',
		}),
		defineType({
			type: 'number',
			name: 'upper_limit_size',
			title: 'Upper Limit Size (sq. ft)',
		}),
		defineType({
			name: 'location',
			title: 'Location',
			type: 'object',
			fields: [
				{type: 'geopoint', name: 'geolocation', title: 'GeoLocation'},
				{type: 'string', name: 'area', title: 'Area'},
				{type: 'string', name: 'city', title: 'City'},
			],
		}),
    ]
}
