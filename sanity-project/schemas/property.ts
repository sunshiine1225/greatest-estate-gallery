import {defineType} from 'sanity'

export default {
	name: 'property',
	type: 'document',
	title: 'Property',
	fields: [
		defineType({
			title: 'Property Name',
			name: 'property_name',
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
			name: 'type',
			title: 'Property Type',
			type: 'string',
			options: {
				list: [
					'1 BHK',
					'1-2 BHK',
					'2 BHK',
					'2-3 BHK',
					'3 BHK',
					'1-3 BHK',
					'4 BHK',
					'2-4 BHK',
					'3-4 BHK',
				],
				layout: 'dropdown',
			},
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
			name: 'full_desc',
			title: 'Description',
			type: 'string',
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
	],
}
