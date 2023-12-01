import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  tabs: [
    { id: 1, name: "Hotels", icon: "icon-bed" },
    { id: 2, name: "Tours", icon: "icon-destination" },
    { id: 3, name: "Activities", icon: "icon-ski" },
    { id: 4, name: "Holiday-Rentals", icon: "icon-home" },
    { id: 5, name: "Cars", icon: "icon-car" },
    { id: 6, name: "Golfs", icon: "icon-yatch" },
    { id: 7, name: "Flights", icon: "icon-tickets" },
  ],
  currentTab: "Hotel",
};

export const findPlaceSlice = createSlice({
  name: "find-place",
  initialState,
  reducers: {
    addCurrentTab: (state, { payload }) => {
      state.currentTab = payload;
    },
  },
});

export const { addCurrentTab } = findPlaceSlice.actions;
export default findPlaceSlice.reducer;
