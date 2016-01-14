package fr.guillaumerose.operators;

import static fr.guillaumerose.Lists.newArrayList;
import static fr.guillaumerose.Numbers.isNumeric;

import java.util.List;

public class Number implements Formula {
	private final Integer value;

	public Number(Integer value) {
		this.value = value;
	}

	@Override
	public List<String> variables() {
		return newArrayList();
	}

	@Override
	public Integer compute(List<Integer> inputs) {
		return value;
	}

	public static Formula parse(String input) {
		if (isNumeric(input)) {
			return new Number(Integer.valueOf(input));
		}
		return new Copy(input);
	}
}