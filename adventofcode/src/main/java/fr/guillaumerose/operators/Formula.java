package fr.guillaumerose.operators;

import java.util.List;

public interface Formula {
	public List<String> variables();

	public Integer compute(List<Integer> values);
}