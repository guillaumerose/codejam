package fr.guillaumerose.operators;

import static fr.guillaumerose.Lists.newArrayList;

import java.util.List;

public class Copy implements Formula {
	private final String variable;

	public Copy(String variable) {
		this.variable = variable;
	}

	@Override
	public List<String> variables() {
		return newArrayList(variable);
	}

	@Override
	public Integer compute(List<Integer> inputs) {
		return inputs.get(0);
	}
}