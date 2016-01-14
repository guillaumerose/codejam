package fr.guillaumerose;

import java.util.ArrayList;
import java.util.List;

public class Lists {
	@SafeVarargs
	public static <T> List<T> newArrayList(T... values) {
		List<T> list = new ArrayList<>();
		for (int i = 0; i < values.length; i++) {
			list.add(values[i]);
		}
		return list;
	}
}